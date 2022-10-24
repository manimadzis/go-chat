package postgres

import (
	"context"
	"go-chat/internal/chat/entities"
	"go-chat/internal/chat/message"
	"go-chat/internal/chat/message/db"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
)

type storage struct {
	db     *gorm.DB
	logger *logging.Logger
}

func (s *storage) AutoMigrate(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&entities.Message{})
}

func (s *storage) Create(ctx context.Context, dto *message.CreateMessageDTO) (*entities.Message, error) {
	msg := entities.Message{
		Text:   dto.Text,
		ChatID: dto.ChatID,
		UserID: dto.UserID,
	}
	err := s.db.WithContext(ctx).Create(&msg).Error
	if err != nil {
		return nil, db.UnknownErr(err)
	}

	return &msg, nil
}

func (s *storage) Delete(ctx context.Context, dto *message.DeleteMessageDTO) error {
	msg := entities.Message{ID: dto.ID}
	return s.db.WithContext(ctx).Delete(&msg).Error
}

func (s *storage) Update(ctx context.Context, dto *message.UpdateMessageDTO) error {
	return s.db.WithContext(ctx).Model(dto.OldMessage).Updates(dto.NewMessage).Error
}

func (s *storage) FindByID(ctx context.Context, id uint) (*entities.Message, error) {
	msg := entities.Message{}
	err := s.db.WithContext(ctx).Take(&msg, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, db.ErrUnknownMessage
		}
		return nil, db.UnknownErr(err)
	}
	return &msg, nil
}

func (s *storage) FindByChat(ctx context.Context, dto *message.FindByChatDTO) ([]entities.Message, error) {
	var msgs []entities.Message
	err := s.db.WithContext(ctx).Offset(dto.Offset).Limit(dto.Limit).Find(&msgs).Error
	if err != nil {
		return nil, db.UnknownErr(err)
	}
	return msgs, nil
}

func NewStorage(gormDB *gorm.DB, logger *logging.Logger) message.Storage {
	return &storage{
		db:     gormDB,
		logger: logger,
	}
}
