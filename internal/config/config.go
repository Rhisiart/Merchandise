package config

import "github.com/kelseyhightower/envconfig"

const (
	prefix = "Merch"
)

type Configuration struct {
	HTTPServer
	Database
}

type HTTPServer struct {
	Port int `envconfig:"PORT" default:8080`
}

type Database struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" required:true`
}

func Load() (Configuration, error) {
	var cfg Configuration
	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
