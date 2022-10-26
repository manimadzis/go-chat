package service

import (
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type Service struct {
	UserService    UserService
	MessageService MessageService
	ChatService    ChatService
}

func New(repo *repository.Repository, logger *logging.Logger) *Service {
	return &Service{
		UserService:    NewUserService(repo, logger),
		MessageService: NewMessageService(repo, logger),
		ChatService:    NewChatService(repo, logger),
	}
}
