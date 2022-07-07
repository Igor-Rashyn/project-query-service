package main

import (
	"github.com/Igor-Rashyn/project-query-service/common/env"
	"github.com/Igor-Rashyn/project-query-service/common/logger"
)

var appConfig *AppConfig

func main() {
	if err := env.New(); err != nil {
		panic(err)
	}
	appConfig, err := NewAppConfig()
	if err != nil {
		panic(err)
	}

	log := logger.New(appConfig.ServiceName, appConfig.Env)
	defer log.Sync()

}
