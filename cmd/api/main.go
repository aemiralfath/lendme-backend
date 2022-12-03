package main

import (
	"final-project-backend/config"
	"final-project-backend/internal/server"
	"final-project-backend/pkg/logger"
	"final-project-backend/pkg/postgres"
	"log"
)

func main() {
	log.Println("Starting api server")
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	gormDB, err := postgres.NewGormDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	}
	appLogger.Infof("Postgres connected")

	s := server.NewServer(cfg, gormDB, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
