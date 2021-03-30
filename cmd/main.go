package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/luckyshmo/api-example/config"
	"github.com/luckyshmo/api-example/pkg/handler"
	"github.com/luckyshmo/api-example/pkg/repository"
	"github.com/luckyshmo/api-example/pkg/repository/pg"
	"github.com/luckyshmo/api-example/pkg/service"
	"github.com/luckyshmo/api-example/server"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	//Mat Ryer advice to handle all app errors
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

//main func
func run() error {
	// config
	cfg := config.Get()

	// logger configuration
	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel) //using debug lvl if we can't parse
		logrus.Warn("Using debug level logger")
	} else {
		logrus.SetLevel(lvl)
	}
	var JSONF = new(logrus.JSONFormatter)
	JSONF.TimestampFormat = time.RFC3339
	logrus.SetFormatter(JSONF)   //todo could be configured via Env
	logrus.SetReportCaller(true) //todo could be configured via Env

	logrus.Print("Run Started")

	//Init DB
	db, err := pg.NewPostgresDB(pg.Config{
		Host:     cfg.PgHOST,
		Port:     cfg.PgPORT,
		Username: "postgres", //TODO config
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "example",
	})
	if err != nil {
		return errors.Wrap(err, "failed to initialize db")
	}

	//Init main components
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//starting server
	srv := new(server.Server) //TODO? server.Server should be *serviceName*.server
	go func() {
		if err := srv.Run(cfg.AppPort, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	return nil
}
