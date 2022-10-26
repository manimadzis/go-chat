package service

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type MessageService interface {
}

type messageService struct {
	repo   *repository.Repository
	logger *logging.Logger
}

func NewMessageService(repo *repository.Repository, logger *logging.Logger) MessageService {
	return &messageService{
		repo:   repo,
		logger: logger,
	}
}

func (m *messageService) Send(ctx context.Context, dto *domain.CreateMessageDTO) error {
	m.logger.Tracef("Started creating the message: %v", dto)
	_, err := m.repo.MessageRepo.Create(ctx, dto)
	if err != nil {
		m.logger.Errorf("Failed sending message: %v", dto)
		return err
	}
	m.logger.Tracef("Successfully created the message: %v", dto)
	return nil
}

func (m *messageService) FindByChat(ctx context.Context, dto *domain.FindMessageByChatDTO) ([]domain.Message, error) {
	m.logger.Tracef("Started looking for messages: %v", dto)
	messages, err := m.repo.MessageRepo.FindByChat(ctx, dto)
	if err != nil {
		m.logger.Errorf("Failed searching messages: %v", dto)
		return []domain.Message{}, err
	}
	m.logger.Tracef("Successfully found messages: %v", dto)
	return messages, nil
}
