//go:build !dev
// +build !dev

package main

import (
	src "github.com/Mirsadikovv/idoctor_bot/app"
	"github.com/Mirsadikovv/shared/sharedutil"
)

func fillEnv(env *src.Env) {
	sharedutil.Load(env)
}
