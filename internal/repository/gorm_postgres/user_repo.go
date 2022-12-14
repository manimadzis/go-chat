package gorm_postgres

import (
	"context"
	"github.com/lib/pq"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func (s *userRepo) FindByRefreshToken(ctx context.Context, refreshToken string) (*domain.User, error) {
	user := domain.User{RefreshToken: refreshToken}
	err := s.db.WithContext(ctx).Where(&user).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repository.ErrUnknownRefreshToken
		}
		return nil, repository.UnknownErr(err)
	}
	return &user, nil
}

func (s *userRepo) SetRefreshToken(ctx context.Context, dto domain.SetRefreshTokenDTO) error {
	return s.db.WithContext(ctx).Model(&domain.User{ID: dto.UserId}).Updates(&domain.User{
		RefreshToken: dto.RefreshToken,
		ExpiredAt:    dto.RefreshTokenExpiredAt,
	}).Error
}

func (s *userRepo) FindById(ctx context.Context, id uint) (*domain.User, error) {
	u := &domain.User{ID: id}
	if err := s.db.WithContext(ctx).Take(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repository.ErrUnknownUser
		}
		return nil, repository.UnknownErr(err)
	}

	return u, nil
}

func (s *userRepo) AutoMigrate(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&domain.User{})
}

func (s *userRepo) Create(ctx context.Context, dto domain.CreateUserDTO) (*domain.User, error) {
	user := &domain.User{
		Login:        dto.Login,
		PasswordHash: dto.Password}

	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code.Name() == "unique_violation" {
			return nil, repository.ErrDuplicatedLogin
		}

		return nil, repository.UnknownErr(err)
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db}
}
