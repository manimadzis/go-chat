package service

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type ChatService struct {
	repo   *repository.Repository
	logger *logging.Logger
}

func NewChatService(repo *repository.Repository, logger *logging.Logger) *ChatService {
	return &ChatService{
		repo:   repo,
		logger: logger,
	}
}

func (c *ChatService) Create(ctx context.Context, dto *domain.CreateChatDTO) error {
	//TODO: implement
	panic("IMPLEMENT")
}
