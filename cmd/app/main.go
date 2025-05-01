package main

import (
	"github.com/vnFuhung2903/rubyams-backend/config"
	postgres "github.com/vnFuhung2903/rubyams-backend/pkg"
)

func main() {
	config := config.NewConfig()
	postgres.InitDatabase(config)
}
