package main

import (
	"git.sriss.uz/mehnat/inspector_platform/src"
)

// @title MyDream API
// @version 1.0
// @description This is a MyDream API.
// @BasePath  /api/v1
// @Schemes https
// @securityDefinitions.apikey ApiKeyAuth
// @BasePath /api/v1
// @in header
// @name Authorization
func main() {
	var env src.Env
	{
		fillEnv(&env)
	}

	src.Exec(&env)
}
