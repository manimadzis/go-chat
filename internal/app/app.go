package app

import (
	"context"
	"go-chat/internal/config"
	repository "go-chat/internal/repository/gorm_postgres"
	dbclient "go-chat/pkg/dbclient/postgres"
	"go-chat/pkg/gormclient"
	"go-chat/pkg/logging"
)

func Run() {
	logger := logging.Get()

	conf, err := config.Load()
	if err != nil {
		logger.Panicf("Can't load config: %v", err)
	}

	dbConn, err := dbclient.New(conf.DB)
	if err != nil {
		logger.Panicf("Cannot connect to DB")
	}
	defer dbConn.Close()

	gormDB, err := gormclient.New(dbConn)
	if err != nil {
		logger.Panicf("Cannot open GORM connection")
	}

	ctx := context.Background()
	repo := repository.New(gormDB, logger)
	err = repo.UserRepo.AutoMigrate(ctx)
	logger.Info(err)
}
