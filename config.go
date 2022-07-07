package main

import "github.com/Igor-Rashyn/project-query-service/common/env"

type AppConfig struct {
	ServiceName string
	Env         string
	Port        int
}

func NewAppConfig() (*AppConfig, error) {
	port, err := env.GetInt("HTTP_PORT", 3000)
	if err != nil {
		return nil, err
	}
	return &AppConfig{
		ServiceName: env.Get("SERVICE_NAME"),
		Env:         env.Get("ENV"),
		Port:        port,
	}, nil
}
