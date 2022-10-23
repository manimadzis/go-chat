package gormclient

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-chat/pkg/dbclient/postgres"
	"gorm.io/gorm"
)

var postgresMock *sqlx.DB = nil

func NewMock() (*gorm.DB, error) {
	var err error
	if postgresMock == nil {
		postgresMock, err = postgres.NewMock()
	}
	if err != nil {
		return nil, fmt.Errorf("can't open connection to db: %v", err)
	}
	return New(postgresMock)
}
