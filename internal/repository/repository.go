package repository

import (
	"context"
	"go-chat/internal/domain"
)

type UserRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *domain.CreateUserDTO) (*domain.User, error)
	FindById(ctx context.Context, id uint) (*domain.User, error)
}

type ChatRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *domain.CreateChatDTO) (*domain.Chat, error)
	Delete(ctx context.Context, dto *domain.DeleteChatDTO) error
	FindByID(ctx context.Context, id uint) (*domain.Chat, error)
}

type MessageRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *domain.CreateMessageDTO) (*domain.Message, error)
	Delete(ctx context.Context, dto *domain.DeleteMessageDTO) error
	Update(ctx context.Context, dto *domain.UpdateMessageDTO) error

	FindByID(ctx context.Context, id uint) (*domain.Message, error)
	FindByChat(ctx context.Context, dto *domain.FindMessageByChatDTO) ([]domain.Message, error)
}

type Repository struct {
	UserRepo    UserRepository
	ChatRepo    ChatRepository
	MessageRepo MessageRepository
}
