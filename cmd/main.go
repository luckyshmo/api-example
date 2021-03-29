package main

import (
	"github.com/luckyshmo/api-example/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Mat Ryer advice to handle all app errors
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

//main func
func run() error {

	// config
	cfg := config.Get()
	logrus.Info(cfg)

	return nil
}
