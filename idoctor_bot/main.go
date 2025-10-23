package main

import (
	cmd "github.com/Mirsadikovv/idoctor_v2/app"
)

func main() {
	var env cmd.Env
	{
		fillEnv(&env)
	}

	cmd.Exec(&env)
}
