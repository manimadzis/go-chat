package gorm_postgres

import (
	"context"
	"github.com/lib/pq"
	"go-chat/internal/entities"
	"go-chat/internal/repository"
	"go-chat/internal/repository/dto"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
)

type chatRepo struct {
	db     *gorm.DB
	logger *logging.Logger
}

func (r *chatRepo) AutoMigrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&entities.Chat{})
}

func (r *chatRepo) Create(ctx context.Context, dto *dto.CreateChatDTO) (*entities.Chat, error) {
	chat := entities.Chat{
		Name: dto.Name,
	}
	err := r.db.WithContext(ctx).Create(&chat).Error
	if err != nil {
		return nil, repository.UnknownErr(err)
	}
	return &chat, nil
}

func (r *chatRepo) Delete(ctx context.Context, dto *dto.DeleteChatDTO) error {
	chat := entities.Chat{
		ID: dto.ID,
	}
	err := r.db.WithContext(ctx).Delete(&chat).Error
	if err != nil {
		return repository.UnknownErr(err)
	}
	return nil
}

func (r *chatRepo) FindByID(ctx context.Context, id uint) (*entities.Chat, error) {
	chat := entities.Chat{
		ID: id,
	}
	err := r.db.WithContext(ctx).Take(&chat).Error
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code.Name() == "unique_violation" {
				return nil, repository.ErrUnknownChat
			}
		}
		return nil, repository.UnknownErr(err)
	}
	return &chat, nil
}

func NewChatRepository(gormDB *gorm.DB, logger *logging.Logger) repository.ChatRepository {
	return &chatRepo{db: gormDB, logger: logger}
}
