package chat

import (
	"context"
	"go-chat/internal/chat/entities"
)

type Storage interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *CreateChatDTO) (*entities.Chat, error)
	Delete(ctx context.Context, dto *DeleteChatDTO) error
	FindByID(ctx context.Context, id uint) (*entities.Chat, error)
}
