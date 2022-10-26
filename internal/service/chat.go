package service

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type ChatService interface {
}

type chatService struct {
	repo   *repository.Repository
	logger *logging.Logger
}

func NewChatService(repo *repository.Repository, logger *logging.Logger) ChatService {
	return &chatService{
		repo:   repo,
		logger: logger,
	}
}

func (c *chatService) Create(ctx context.Context, dto *domain.CreateChatDTO) error {
	//TODO: implement
	panic("IMPLEMENT")
}
