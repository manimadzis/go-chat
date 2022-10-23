package postgres

import (
	"context"
	"github.com/lib/pq"
	"go-chat/internal/chat/entities"
	"go-chat/internal/chat/user"
	"go-chat/internal/chat/user/db"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func (s *storage) FindById(ctx context.Context, id uint) (*entities.User, error) {
	u := &entities.User{}
	if err := s.db.WithContext(ctx).First(&u).Error; err != nil {
		//if pqErr, ok := err.(*pq.Error); ok {
		//	pqErr.
		//}
		return nil, err
	}

	return u, nil
}

func (s *storage) AutoMigrate(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&entities.User{})
}

func (s *storage) Create(ctx context.Context, dto *user.CreateUserDTO) (*entities.User, error) {
	user := &entities.User{
		Login:        dto.Login,
		PasswordHash: dto.Password}

	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		pqErr, ok := err.(*pq.Error)
		if !ok {
			return nil, db.ErrUnknown
		}

		if pqErr.Code.Name() == "unique_violation" {
			return nil, db.ErrDuplicatedLogin
		}
	}

	return user, nil
}

func NewRepository(db *gorm.DB) user.Storage {
	return &storage{db}
}
