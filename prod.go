//go:build !dev
// +build !dev

package main

import (
	"git.sriss.uz/shared/shared_service/sharedutil"
	src "github.com/Mirsadikovv/idoctor_bot/app"
)

func fillEnv(env *src.Env) {
	sharedutil.Load(env)
}
