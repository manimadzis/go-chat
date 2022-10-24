package repository

import (
	"context"
	"go-chat/internal/entities"
	"go-chat/internal/repository/dto"
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
	UserRepo    UserRepository
	ChatRepo    ChatRepository
	MessageRepo MessageRepository
}
