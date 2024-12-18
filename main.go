package main

import (
	"github.com/isaacgoncod/go-project-example/config"
	"github.com/isaacgoncod/go-project-example/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errf("Config Initialization Error: %v", err)
		return
	}

	router.Initialize()
}
