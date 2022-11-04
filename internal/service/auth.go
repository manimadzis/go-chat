package service

import (
	"context"
	"fmt"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/jwtmanager"
	"go-chat/pkg/logging"
)

type authService struct {
	jwtManager jwtmanager.JWTManager
	logger     *logging.Logger
	repo       *repository.Repository
}

type AuthService interface {
}

func NewAuthService(repo *repository.Repository, logger *logging.Logger) (AuthService, error) {
	if logger == nil {
		return nil, fmt.Errorf("nil logger")
	}

	if repo == nil {
		return nil, fmt.Errorf("nil repo")
	}

	var err error
	service := authService{logger: logger}

	service.jwtManager, err = jwtmanager.New(jwtmanager.Config{SignKey: "123"})
	if err != nil {
		fmt.Errorf("can't create AuthService: %v", err)
	}

	return &service, nil
}

func (a *authService) CreateUserSession(ctx context.Context, userId uint) (*jwtmanager.JWTAuth, error) {
	return a.newSession(ctx, userId)
}
func (a *authService) RefreshUserSession(ctx context.Context, refreshToken string) (*jwtmanager.JWTAuth, error) {
	return a.refreshSession(ctx, refreshToken)
}

func (a *authService) refreshSession(ctx context.Context, refreshToken string) (*jwtmanager.JWTAuth, error) {
	a.logger.Tracef("Refreshing JWT for refresh token %s", refreshToken)

	user, err := a.repo.UserRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		a.logger.Error("Can't find user  by refresh token: %v", err)
		return nil, err
	}

	newAuth, err := a.jwtManager.RefreshJWT(fmt.Sprintf("%d", user.ID), refreshToken)
	if err != nil {
		a.logger.Error("Can't refresh token: %v", err)
		return nil, err
	}

	return newAuth, nil
}

func (a *authService) newSession(ctx context.Context, userId uint) (*jwtmanager.JWTAuth, error) {
	a.logger.Tracef("Create new session for user %d", userId)
	jwtAuth, err := a.jwtManager.NewJWTAuth(fmt.Sprintf("%d", userId))
	if err != nil {
		return nil, fmt.Errorf("can't create new session: %v", err)
	}

	err = a.repo.UserRepo.SetRefreshToken(ctx, domain.SetRefreshTokenDTO{
		UserId:                userId,
		RefreshToken:          jwtAuth.RefreshToken,
		RefreshTokenExpiredAt: jwtAuth.RefreshTokenExpiredAt,
	})
	if err != nil {
		return nil, fmt.Errorf("can't create new session in db: %v", err)
	}
	return jwtAuth, nil
}
