package app

import (
	"fmt"
	"go-chat/internal/config"
	"go-chat/internal/server"
	"go-chat/pkg/logging"
)

func Run() {
	logger := logging.Get()

	conf, err := config.Load()
	if err != nil {
		logger.Panicf("Can't load config: %v", err)
	}

	fmt.Printf("%#v", conf)
	//
	//dbConn, err := dbclient.New(conf.DB)
	//if err != nil {
	//	logger.Panicf("Cannot connect to DB")
	//}
	//defer dbConn.Close()
	//
	//gormDB, err := gormclient.New(dbConn)
	//if err != nil {
	//	logger.Panicf("Cannot open GORM connection")
	//}
	//
	//ctx := context.Background()
	//repo := repository.New(gormDB, logger)
	//err = repo.UserRepo.AutoMigrate(ctx)
	//logger.Info(err)

	s := server.New(conf.Server, logging.Get())
	s.Start()
}
