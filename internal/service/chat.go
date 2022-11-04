package service

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type ChatService interface {
	Create(ctx context.Context, dto domain.CreateChatDTO) error
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

func (c *chatService) Create(ctx context.Context, dto domain.CreateChatDTO) error {
	c.logger.Trace("Creating new chat: %#v", dto)
	_, err := c.repo.ChatRepo.Create(ctx, dto)
	if err != nil {
		c.logger.Errorf("Can't create new chat: %v", err)
	}
	return err
}

func (c *chatService) Update(ctx context.Context, dto domain.UpdateChatDTO) error {
	c.logger.Trace("Update chat: %#v", dto)
	return c.repo.ChatRepo.Update(ctx, dto)
}
