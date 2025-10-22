//go:build dev
// +build dev

package main

import (
	"github.com/Mirsadikovv/idoctor_platform/src"

	"github.com/Mirsadikovv/shared/sharedutil"
)

// @title github.com/Mirsadikovv/idoctor_platform API
// @version 1.0
// @description This is a github.com/Mirsadikovv/idoctor_platform API.
// @BasePath  /api/v1
// @Schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func fillEnv(env *src.Env) {
	sharedutil.MustLoad(env, ".env.dev")
}
