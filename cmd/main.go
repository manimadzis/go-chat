package main

import (
	"context"
	userDB "go-chat/internal/chat/user/db/postgres"
	dbclient "go-chat/pkg/dbclient/postgres"
	"go-chat/pkg/gormclient"
	"go-chat/pkg/logging"
)

func main() {
	logger := logging.Get()

	dbConfig := dbclient.Config{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "pass",
		Database: "chat",
	}

	dbConn, err := dbclient.New(dbConfig)
	if err != nil {
		logger.Panicf("Cannot connect to DB")
	}
	defer dbConn.Close()

	gormDB, err := gormclient.New(dbConn)
	if err != nil {
		logger.Panicf("Cannot open GORM connection")
	}

	ctx := context.Background()
	repo := userDB.NewRepository(gormDB)
	err = repo.AutoMigrate(ctx)
	logger.Error(err)

	//_, err = repo.Create(ctx, &user.CreateUserDTO{Login: "123", Password: "123"})
	//logger.Error(err)

}
