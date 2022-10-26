package app

import (
	"go-chat/internal/config"
	repository "go-chat/internal/repository/gorm_postgres"
	"go-chat/internal/server"
	"go-chat/internal/service"
	dbclient "go-chat/pkg/dbclient/postgres"
	"go-chat/pkg/gormclient"
	"go-chat/pkg/logging"
	"log"
)

func Run() {
	conf, err := config.Load()
	if err != nil {
		log.Panicf("Can't load config: %v", err)
	}
	log.Println("Load config")

	logger := logging.Get()
	if logger == nil {
		log.Panicf("Can't create logger: %v", err)
	}
	log.Println("Init logger")

	logger.Info("Creating connection to DB...")
	dbConn, err := dbclient.New(dbclient.Config{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		Username: conf.DB.Username,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
	})
	if err != nil {
		logger.Panicf("Cannot connect to DB")
	}
	defer dbConn.Close()
	logger.Info("Create connection to DB")

	logger.Info("Opening gorm connection to gorm...")
	gormDB, err := gormclient.New(dbConn)
	if err != nil {
		logger.Panicf("Cannot open GORM connection")
	}
	logger.Info("Open gorm connection")

	repo := repository.New(gormDB, logger)

	s := server.New(&server.Config{
		Host: conf.Server.Host,
		Port: conf.Server.Port,
	}, service.New(repo, logger), logger)
	logger.Info("Starting app...")
	logger.Error(s.Start())
}
