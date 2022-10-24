package service

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type UserService struct {
	repo   *repository.Repository
	logger *logging.Logger
}

func NewUserService(repo *repository.Repository, logger *logging.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) SingUp(ctx context.Context, dto *domain.CreateUserDTO) error {
	u.logger.Tracef("Start to sign up new user: %v", dto)

	_, err := u.repo.UserRepo.Create(ctx, dto)
	if err != nil {
		u.logger.Errorf("User creation is failed: %v", err)
		return err
	}
	u.logger.Tracef("Successfully signed up new user: %v", dto)
	return nil
}
