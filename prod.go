//go:build !dev
// +build !dev

package main

import (
	"github.com/Mirsadikovv/idoctor_platform/src"

	"github.com/Mirsadikovv/shared/sharedutil"
)

func fillEnv(env *src.Env) {
	sharedutil.Load(env)
}
