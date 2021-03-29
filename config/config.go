package config

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
)

// Config. Should be filled from Env. Use launch.json(vscode) on local machine
type Config struct {
	LogLevel         string `envconfig:"LOG_LEVEL"`
	PgHOST           string `envconfig:"PG_HOST"`
	PgPORT           string `envconfig:"PG_PORT"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	AppPort          string `envconfig:"APP_PORT"`
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
