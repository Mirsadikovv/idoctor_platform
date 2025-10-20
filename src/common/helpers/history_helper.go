package helpers

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterHistoryCallbacks - Create, Update, Delete callbacklarni register qiladi
func RegisterHistoryCallbacks(db *gorm.DB, e *echo.Echo) {
	db.Callback().Create().Before("gorm:create").Register("history:set_local", func(tx *gorm.DB) {
		setLocalVars(tx)
	})
	db.Callback().Update().Before("gorm:update").Register("history:set_local", func(tx *gorm.DB) {
		setLocalVars(tx)
	})
	db.Callback().Delete().Before("gorm:delete").Register("history:set_local", func(tx *gorm.DB) {
		setLocalVars(tx)
	})
}

func setLocalVars(tx *gorm.DB) {
	ctxVal := tx.Statement.Context.Value("history_context")
	if ctxVal == nil {
		return
	}

	data, ok := ctxVal.(map[string]string)
	if !ok {
		return
	}

	userID := data["user_id"]
	ip := data["ip"]
	path := data["path"]
	method := data["method"]

	if userID == "" {
		userID = "0"
	}
	if ip == "" {
		ip = "NULL"
	} else {
		ip = fmt.Sprintf("'%s'", ip)
	}
	if path == "" {
		path = "NULL"
	} else {
		path = fmt.Sprintf("'%s'", path)
	}
	if method == "" {
		method = "NULL"
	} else {
		method = fmt.Sprintf("'%s'", method)
	}

	setVars := []string{
		fmt.Sprintf("SET app.current_user_id = %s", userID),
		fmt.Sprintf("SET app.current_request_ip = %s", ip),
		fmt.Sprintf("SET app.current_request_path = %s", path),
		fmt.Sprintf("SET app.current_request_method = %s", method),
	}

	for _, sql := range setVars {
		if err := tx.Exec(sql).Error; err != nil {
			log.Println("⚠️ Failed to set variable:", sql, "error:", err)
		}
	}
}
