package postgres

import (
	"context"
	"github.com/lib/pq"
	"go-chat/internal/chat/chat"
	"go-chat/internal/chat/chat/db"
	"go-chat/internal/chat/entities"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
)

type storage struct {
	db     *gorm.DB
	logger *logging.Logger
}

func (r *storage) AutoMigrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&entities.Chat{})
}

func (r *storage) Create(ctx context.Context, dto *chat.CreateChatDTO) (*entities.Chat, error) {
	chat := entities.Chat{
		Name: dto.Name,
	}
	err := r.db.WithContext(ctx).Create(&chat).Error
	if err != nil {
		return nil, db.UnknownErr(err)
	}
	return &chat, nil
}

func (r *storage) Delete(ctx context.Context, dto *chat.DeleteChatDTO) error {
	chat := entities.Chat{
		ID: dto.ID,
	}
	err := r.db.WithContext(ctx).Delete(&chat).Error
	if err != nil {
		return db.UnknownErr(err)
	}
	return nil
}

func (r *storage) FindByID(ctx context.Context, id uint) (*entities.Chat, error) {
	chat := entities.Chat{
		ID: id,
	}
	err := r.db.WithContext(ctx).Take(&chat).Error
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code.Name() == "unique_violation" {
				return nil, db.ErrUnknownChat
			}
		}
		return nil, db.UnknownErr(err)
	}
	return &chat, nil
}

func NewStorage(gormDB *gorm.DB, logger *logging.Logger) chat.Storage {
	return &storage{db: gormDB, logger: logger}
}
