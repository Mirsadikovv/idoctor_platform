package src

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"path"
	"strings"

	auth_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service"
	auth_dto "github.com/Mirsadikovv/idoctor_platform/src/module/auth_service/dto"
	bot_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/bot_service"
	bot_model "github.com/Mirsadikovv/idoctor_platform/src/module/bot_service/model"
	file_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/file_service"
	file_model "github.com/Mirsadikovv/idoctor_platform/src/module/file_service/model"
	language_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/language_service"
	language_model "github.com/Mirsadikovv/idoctor_platform/src/module/language_service/model"
	organization_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service"
	organization_model "github.com/Mirsadikovv/idoctor_platform/src/module/organization_service/model"
	role_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/role_service"
	role_model "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/model"
	user_cmd "github.com/Mirsadikovv/idoctor_platform/src/module/user_service"
	user_model "github.com/Mirsadikovv/idoctor_platform/src/module/user_service/model"

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

	authMiddleware := middleware.NewAuthEchoMiddleware[*auth_dto.AuthUser](jwtConfig, memoryCache, db)

	auth_cmd.Cmd(router, db, log, authMiddleware)
	bot_cmd.Cmd(router, db, log, authMiddleware)
	file_cmd.Cmd(router, db, log, authMiddleware)
	language_cmd.Cmd(router, db, log, authMiddleware)
	organization_cmd.Cmd(router, db, log, authMiddleware)
	role_cmd.Cmd(router, db, log, authMiddleware)
	user_cmd.Cmd(router, db, log, authMiddleware)

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

	err := db.AutoMigrate(models...)

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

func Seed(db *gorm.DB) {
	// Seed functions for appeal, setting models have been removed
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
