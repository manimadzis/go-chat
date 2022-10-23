package user

import (
	"context"
	"go-chat/internal/chat/entities"
)

type Service interface {
	Create(ctx context.Context, dto *CreateUserDTO) *entities.User
}
