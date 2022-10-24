package gorm_postgres

import (
	"context"
	"github.com/lib/pq"
	"go-chat/internal/entities"
	"go-chat/internal/repository"
	"go-chat/internal/repository/dto"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func (s *userRepo) FindById(ctx context.Context, id uint) (*entities.User, error) {
	u := &entities.User{ID: id}
	if err := s.db.WithContext(ctx).Take(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repository.ErrUnknownUser
		}
		return nil, repository.UnknownErr(err)
	}

	return u, nil
}

func (s *userRepo) AutoMigrate(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&entities.User{})
}

func (s *userRepo) Create(ctx context.Context, dto *dto.CreateUserDTO) (*entities.User, error) {
	user := &entities.User{
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
