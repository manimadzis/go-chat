package gormclient

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	gorm_pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func New(db *sqlx.DB) (*gorm.DB, error) {
	dbName := db.DriverName()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	switch dbName {
	case "postgres":
		return gorm.Open(gorm_pg.New(gorm_pg.Config{Conn: db}), &gorm.Config{Logger: newLogger})
	default:
		panic(fmt.Errorf("cannot use %s DB", dbName))
	}
}
