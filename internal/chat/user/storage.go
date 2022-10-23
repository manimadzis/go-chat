package user

import (
	"context"
	"go-chat/internal/chat/entities"
)

type Storage interface {
	AutoMigrate(ctx context.Context) error
	Create(ctx context.Context, dto *CreateUserDTO) (*entities.User, error)
	FindById(ctx context.Context, id uint) (*entities.User, error)
}
