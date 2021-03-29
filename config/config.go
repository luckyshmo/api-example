package config

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
)

// Config is a config :)
type Config struct {
	LogLevel         string `envconfig:"LOG_LEVEL"`
	PgURL            string `envconfig:"PG_URL"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			logrus.Fatal(err)
		}
	})
	return &config
}
