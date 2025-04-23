package main

import (
	"github.com/vnFuhung2903/rubyams-backend/config"
	"github.com/vnFuhung2903/rubyams-backend/env"
)

func main() {
	env := env.InitEnv()
	config.ConnectPostgresDb(env)
}
