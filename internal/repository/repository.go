package repository

import (
	"context"
	"go-chat/internal/entities"
	"go-chat/internal/repository/dto"
	"go-chat/internal/repository/gorm_postgres"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
)

type UserRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *dto.CreateUserDTO) (*entities.User, error)
	FindById(ctx context.Context, id uint) (*entities.User, error)
}

type ChatRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *dto.CreateChatDTO) (*entities.Chat, error)
	Delete(ctx context.Context, dto *dto.DeleteChatDTO) error
	FindByID(ctx context.Context, id uint) (*entities.Chat, error)
}

type MessageRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *dto.CreateMessageDTO) (*entities.Message, error)
	Delete(ctx context.Context, dto *dto.DeleteMessageDTO) error
	Update(ctx context.Context, dto *dto.UpdateMessageDTO) error

	FindByID(ctx context.Context, id uint) (*entities.Message, error)
	FindByChat(ctx context.Context, dto *dto.FindByChatDTO) ([]entities.Message, error)
}

type Repository struct {
	UserStorage    UserRepository
	ChatStorage    ChatRepository
	MessageStorage MessageRepository
}

func New(gormDB *gorm.DB, logger *logging.Logger) *Repository {
	return &Repository{
		UserStorage:    gorm_postgres.NewUserRepository(gormDB),
		ChatStorage:    gorm_postgres.NewChatRepository(gormDB, logger),
		MessageStorage: gorm_postgres.NewMessageRepository(gormDB, logger),
	}
}
