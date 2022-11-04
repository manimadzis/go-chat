package repository

import (
	"context"
	"go-chat/internal/domain"
)

type UserRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto domain.CreateUserDTO) (*domain.User, error)
	SetRefreshToken(ctx context.Context, dto domain.SetRefreshTokenDTO) error

	FindById(ctx context.Context, id uint) (*domain.User, error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.User, error)
}

type ChatRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto domain.CreateChatDTO) (*domain.Chat, error)
	Update(ctx context.Context, dto domain.UpdateChatDTO) error
	Delete(ctx context.Context, dto domain.DeleteChatDTO) error

	FindByID(ctx context.Context, id uint) (*domain.Chat, error)
}

type MessageRepository interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto domain.CreateMessageDTO) (*domain.Message, error)
	Update(ctx context.Context, dto domain.UpdateMessageDTO) error
	Delete(ctx context.Context, dto domain.DeleteMessageDTO) error

	FindByID(ctx context.Context, id uint) (*domain.Message, error)
	FindByChat(ctx context.Context, dto domain.FindMessageByChatDTO) ([]domain.Message, error)
}

type Repository struct {
	UserRepo    UserRepository
	ChatRepo    ChatRepository
	MessageRepo MessageRepository
}
