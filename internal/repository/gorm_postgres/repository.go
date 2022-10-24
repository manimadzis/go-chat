package gorm_postgres

import (
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
)

func New(gormDB *gorm.DB, logger *logging.Logger) *repository.Repository {
	return &repository.Repository{
		UserRepo:    NewUserRepository(gormDB),
		ChatRepo:    NewChatRepository(gormDB, logger),
		MessageRepo: NewMessageRepository(gormDB, logger),
	}
}
