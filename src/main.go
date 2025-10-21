package src

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"path"
	"strings"

	"github.com/Mirsadikovv/idoctor_platform/src/common/helpers"
	bot_model "github.com/Mirsadikovv/idoctor_platform/src/module/bot_service/model"
	log_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/log_service"
	log_model "github.com/Mirsadikovv/idoctor_platform/src/module/log_service/model"
	organization_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service"
	organization_model "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/model"
	role_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/role_service"
	role_model "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/model"


	"github.com/Mirsadikovv/shared/jwt"
	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/middleware"
	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/redis"
	"github.com/Mirsadikovv/shared/sharedutil"
	"github.com/Mirsadikovv/shared/swagger"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Env struct {
	HTTP_Host                 string `env:"HTTP_HOST" default:"localhost"`
	HTTP_Port                 int    `env:"HTTP_PORT" default:"80"`
	POSTGRES_Host             string `env:"POSTGRES_HOST"`
	POSTGRES_User             string `env:"POSTGRES_USER"`
	POSTGRES_Password         string `env:"POSTGRES_PASSWORD"`
	POSTGRES_DBName           string `env:"POSTGRES_DB"`
	POSTGRES_Port             int    `env:"POSTGRES_PORT"`
	POSTGRES_SSLMode          string `env:"POSTGRES_SSL_MODE" default:"disable"`
	POSTGRES_TimeZone         string `env:"POSTGRES_TIME_ZONE" default:"UTC"`
	JWT_Secret                string `env:"JWT_SECRET"`
	JWT_Expired               int64  `env:"JWT_EXPIRED"`
	JWT_RefreshExpired        int64  `env:"JWT_REFRESH_EXPIRED"`
	REDIS_Addr                string `env:"REDIS_ADDR"`
	OneIdBaseUrl              string `env:"ONEID_BASE_URL"`
	OneIdClientId             string `env:"ONEID_CLIENT_ID"`
	OneIdClientSecret         string `env:"ONEID_CLIENT_SECRET"`
	OneIdClientRedirectUrl    string `env:"ONEID_CLIENT_REDIRECT_URL"`
	OneIdDashboardRedirectUrl string `env:"ONEID_DASHBOARD_REDIRECT_URL"`
}

func Exec(env *Env) {
	redisConfig := redis.Config{
		Addr: env.REDIS_Addr,
	}

	jwtConfig := &jwt.JwtConfig{
		Secret:         env.JWT_Secret,
		Expired:        env.JWT_Expired,
		RefreshExpired: env.JWT_RefreshExpired,
	}

	gormConfig := &pg.GormConfig{
		SkipDefaultTransaction: true,
	}

	pgConfig := pg.ConnectionConfig{
		Host:     env.POSTGRES_Host,
		User:     env.POSTGRES_User,
		Password: env.POSTGRES_Password,
		DBName:   env.POSTGRES_DBName,
		Port:     env.POSTGRES_Port,
		SSLMode:  env.POSTGRES_SSLMode,
		TimeZone: env.POSTGRES_TimeZone,
	}

	db := pg.Primary(gormConfig, pgConfig)

	if err := migration(db); err != nil {
		log.Println(err)
	}

	//TODO: remove
	// createOrg(db)

	router := echo.New()

	router.Use(echo_middleware.CORS())

	log := logger.New()

	memoryCache := redis.Instance(redisConfig)

	router.Use(helpers.HistoryMiddleware(jwtConfig))

	helpers.RegisterHistoryCallbacks(db, router)

	authMiddleware := middleware.NewAuthEchoMiddleware[*auth_dto.AuthUser](jwtConfig, memoryCache, db)
	{
		oneIdConfig := &auth_dto.OneIdConfig{
			OneIdBaseUrl:              env.OneIdBaseUrl,
			OneIdClientId:             env.OneIdClientId,
			OneIdClientSecret:         env.OneIdClientSecret,
			OneIdClientRedirectUrl:    env.OneIdClientRedirectUrl,
			OneIdDashboardRedirectUrl: env.OneIdDashboardRedirectUrl,
		}

		role_cmd.Cmd(router, db, log, authMiddleware)
		language_cmd.Cmd(router, db, log, authMiddleware)
		file_cmd.Cmd(router, db, log, authMiddleware)

	router.GET("/swagger/dir", swaggerDirs())
	router.GET("/swagger/*", swaggerHandler())

	createOrg(db, router)

	router.Start(fmt.Sprintf(":%d", env.HTTP_Port))

}

func migration(db *gorm.DB) error {

	models := []interface{}{
		&user_model.User{},
		&role_model.Role{},
		&language_model.Language{},
		&language_model.LanguageContent{},
		&file_model.File{},

		&organization_model.Organization{},
		&organization_model.OrganizationTranslation{},
		&organization_model.Invitation{},

		&bot_model.BotUser{},
		&bot_model.BotMessage{},

	}

	models_migration := append(models, &log_model.Log{})

	err := db.AutoMigrate(models_migration...)

	CreateHistoryTriggers(db, models)

	Seed(db)
	if err != nil {
		return err
	}

	return db.Exec(`

        CREATE EXTENSION IF NOT EXISTS pgcrypto;

        CREATE OR REPLACE FUNCTION HASH_MAKE(password TEXT) RETURNS TEXT AS $$
			BEGIN
				RETURN crypt(password, gen_salt('bf'));
			END;
        $$ LANGUAGE plpgsql;

        CREATE OR REPLACE FUNCTION HASH_CHECK(password TEXT, hashed_password TEXT) RETURNS BOOLEAN AS $$
			BEGIN
				RETURN crypt(password, hashed_password) = hashed_password;
			END;
        $$ LANGUAGE plpgsql;
		
	`).Error

}

func CreateHistoryTriggers(DB *gorm.DB, models []interface{}) {

	if err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS histories (
			id SERIAL PRIMARY KEY,
			user_id BIGINT NULL,
			table_name VARCHAR(100),
			row_id BIGINT NULL,
			action VARCHAR(20),
			ip VARCHAR(45) NULL,
			method VARCHAR(10) NULL,
			api TEXT NULL,
			old_value JSONB,
			new_value JSONB,
			created_at TIMESTAMP DEFAULT now()
		);
	`).Error; err != nil {
		log.Fatal("❌ Failed to create histories table: ", err)
	}

	if err := DB.Exec(`
		DO $$
		BEGIN
			CREATE OR REPLACE FUNCTION log_history()
			RETURNS TRIGGER AS $func$
			DECLARE
				v_old JSONB := '{}'::JSONB;
				v_new JSONB := '{}'::JSONB;
				v_uid_text TEXT;
				v_uid BIGINT;
				v_ip TEXT;
				v_path TEXT;
				v_method TEXT;
				v_row_id BIGINT;
				key TEXT;
				old_row JSONB;
				new_row JSONB;
			BEGIN
				v_uid_text := current_setting('app.current_user_id', true);
				IF v_uid_text IS NULL OR v_uid_text = '' THEN
					v_uid := NULL;
				ELSE
					v_uid := v_uid_text::BIGINT;
				END IF;

				v_ip := current_setting('app.current_request_ip', true);
				v_path := current_setting('app.current_request_path', true);
				v_method := current_setting('app.current_request_method', true);

				IF (TG_OP = 'INSERT') THEN
					v_new := to_jsonb(NEW);
					BEGIN v_row_id := NEW.id; EXCEPTION WHEN others THEN v_row_id := NULL; END;
					INSERT INTO histories (user_id, table_name, row_id, action, new_value, ip, api, method, created_at)
					VALUES (v_uid, TG_TABLE_NAME, v_row_id, 'INSERT', v_new, v_ip, v_path, v_method, now());
					RETURN NEW;

				ELSIF (TG_OP = 'UPDATE') THEN
					old_row := to_jsonb(OLD);
					new_row := to_jsonb(NEW);

					-- faqat o‘zgargan maydonlarni yozamiz
					FOR key IN SELECT jsonb_object_keys(new_row) LOOP
						IF old_row->>key IS DISTINCT FROM new_row->>key THEN
							v_old := v_old || jsonb_build_object(key, old_row->key);
							v_new := v_new || jsonb_build_object(key, new_row->key);
						END IF;
					END LOOP;

					BEGIN v_row_id := NEW.id; EXCEPTION WHEN others THEN v_row_id := NULL; END;
					INSERT INTO histories (user_id, table_name, row_id, action, old_value, new_value, ip, api, method, created_at)
					VALUES (v_uid, TG_TABLE_NAME, v_row_id, 'UPDATE', v_old, v_new, v_ip, v_path, v_method, now());
					RETURN NEW;

				ELSIF (TG_OP = 'DELETE') THEN
					v_old := to_jsonb(OLD);
					BEGIN v_row_id := OLD.id; EXCEPTION WHEN others THEN v_row_id := NULL; END;
					INSERT INTO histories (user_id, table_name, row_id, action, old_value, ip, api, method, created_at)
					VALUES (v_uid, TG_TABLE_NAME, v_row_id, 'DELETE', v_old, v_ip, v_path, v_method, now());
					RETURN OLD;
				END IF;

				RETURN NULL;
			END;
			$func$ LANGUAGE plpgsql;
		END $$;
	`).Error; err != nil {
		log.Fatal("❌ Failed to create or replace log_history function: ", err)
	}

	for _, model := range models {
		stmt := &gorm.Statement{DB: DB}
		if err := stmt.Parse(model); err != nil {
			log.Fatalf("❌ Failed to parse model: %v", err)
		}

		table := stmt.Schema.Table
		if strings.TrimSpace(table) == "" {
			log.Printf("⚠️ Skip: model has no table name")
			continue
		}

		triggerName := fmt.Sprintf("%s_history", table)

		sql := fmt.Sprintf(`
			DO $$
			BEGIN
				-- eski triggerni o‘chiramiz
				IF EXISTS (SELECT 1 FROM pg_trigger WHERE tgname = '%s') THEN
					EXECUTE format('DROP TRIGGER IF EXISTS %s ON %s;', '%s', '%s');
				END IF;

				-- yangisini yaratamiz
				EXECUTE format('CREATE TRIGGER %s AFTER INSERT OR UPDATE OR DELETE ON %s FOR EACH ROW EXECUTE FUNCTION log_history();', '%s', '%s');
			END $$;
		`, triggerName, triggerName, table, triggerName, table, triggerName, table)

		if err := DB.Exec(sql).Error; err != nil {
			log.Printf("❌ Failed to (re)create trigger for %s: %v", table, err)
		} else {
			log.Printf("✅ Trigger updated for table: %s", table)
		}
	}
}

func Seed(db *gorm.DB) {

	appeal_model.SeedComplaintTypes(db)
	appeal_model.SeedIndustries(db)
	setting_model.SeedSiteSetting(db)
	appeal_model.SeedAppealStatuses(db)
}

func createOrg(db *gorm.DB, r *echo.Echo) {
	var routers = make(sharedutil.JsonObject)

	for _, route := range r.Routes() {
		if route.Method == "echo_route_not_found" {
			fmt.Println("Skipping route:", route.Path, "Method:", route.Method)
		}
		routers[route.Path] = []string{route.Method}
	}
	fmt.Println(db.Where("roles.id = 1").Delete(&role_model.Role{}).Error)

	db.Create(&role_model.Role{
		ID:          1,
		Name:        "Admin",
		Description: "Admin",
		Permissions: routers,
	})
}

func createOrg_(db *gorm.DB) {

	db.Where("1 = 1").Delete(&organization_model.Organization{})
	db.Where("1 = 1").Delete(&organization_model.Invitation{})

	org := organization_model.Organization{
		SoatoId: 17,
		OrgRoles: pq.Int64Array{
			1,
		},
		ReviewRoleId: 1,
		ParentId:     0,
		Id:           1,
	}

	fmt.Println(db.Create(&org).Error)

	db.Create(&role_model.Role{
		ID:          1,
		Name:        "Admin",
		Description: "Admin",
		Pages: sharedutil.JsonObject{
			"ADMINISTRATION_PAGE": []string{"GET", "POST", "PATCH", "DELETE"},
			"BUTCHERY_PAGE":       []string{"GET", "POST", "PATCH", "DELETE"},
			"BUTCHER_CREATE":      []string{"GET", "POST", "PATCH", "DELETE"},
			"BUTCHER_DELETE":      []string{"GET", "POST", "PATCH", "DELETE"},
		},
	})

	invitation := organization_model.Invitation{
		OrganizationId: org.Id,
		Pin:            52605046520045,
		RoleId:         1,
	}

	fmt.Println(db.Create(&invitation).Error)
}

//go:embed docs
var docs embed.FS

func swaggerHandler() echo.HandlerFunc {
	return swagger.NewSwaggerHandler(docs)
}

func swaggerDirs() echo.HandlerFunc {

	type SwaggerDirs struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	var urls []SwaggerDirs

	fs.WalkDir(docs, ".", func(sysPath string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			return nil
		}
		paths := strings.Split(sysPath, "/")
		serviceName := paths[len(paths)-2]

		urls = append(urls, SwaggerDirs{
			Name: serviceName,
			URL:  path.Join("swagger", serviceName),
		})

		return nil
	})

	return func(c echo.Context) error {
		return c.JSON(200, urls)
	}
}
