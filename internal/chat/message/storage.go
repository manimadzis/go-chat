package message

import (
	"context"
	"go-chat/internal/chat/entities"
)

type Storage interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *CreateMessageDTO) (*entities.Message, error)
	Delete(ctx context.Context, dto *DeleteMessageDTO) error
	Update(ctx context.Context, dto *UpdateMessageDTO) error

	FindByID(ctx context.Context, id uint) (*entities.Message, error)
	FindByChat(ctx context.Context, dto *FindByChatDTO) ([]entities.Message, error)
}
