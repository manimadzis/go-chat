package service

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
)

type UserService interface {
	SignUp(ctx context.Context, dto domain.CreateUserDTO) error
}

type userService struct {
	repo   *repository.Repository
	logger *logging.Logger
}

func NewUserService(repo *repository.Repository, logger *logging.Logger) UserService {
	return &userService{
		repo:   repo,
		logger: logger,
	}
}

func (u *userService) SignUp(ctx context.Context, dto domain.CreateUserDTO) error {
	u.logger.Tracef("Start to sign up new user: %v", dto)

	_, err := u.repo.UserRepo.Create(ctx, dto)
	if err != nil {
		u.logger.Errorf("User creation is failed: %v", err)
		return err
	}
	u.logger.Tracef("Successfully signed up new user: %v", dto)
	return nil
}

func (u *userService) SignIn(ctx context.Context, dto domain.CreateUserDTO) error {
	u.logger.Tracef("Start to sign in user: %v", dto)

	_, err := u.repo.UserRepo.Create(ctx, dto)
	if err != nil {
		u.logger.Errorf("User creation is failed: %v", err)
		return err
	}
	u.logger.Tracef("Successfully signed up new user: %v", dto)
	return nil
}
