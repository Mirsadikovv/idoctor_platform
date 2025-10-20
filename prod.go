//go:build !dev
// +build !dev

package main

import (
	"git.sriss.uz/mehnat/inspector_platform/src"

	"git.sriss.uz/shared/shared_service/sharedutil"
)

func fillEnv(env *src.Env) {
	sharedutil.Load(env)
}
