package gorm_postgres

import (
	"context"
	"go-chat/internal/domain"
	"go-chat/internal/repository"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
)

type messageRepo struct {
	db     *gorm.DB
	logger *logging.Logger
}

func (s *messageRepo) AutoMigrate(ctx context.Context) error {
	return s.db.WithContext(ctx).AutoMigrate(&domain.Message{})
}

func (s *messageRepo) Create(ctx context.Context, dto *domain.CreateMessageDTO) (*domain.Message, error) {
	msg := domain.Message{
		Text:   dto.Text,
		ChatID: dto.ChatID,
		UserID: dto.UserID,
	}
	err := s.db.WithContext(ctx).Create(&msg).Error
	if err != nil {
		return nil, repository.UnknownErr(err)
	}

	return &msg, nil
}

func (s *messageRepo) Delete(ctx context.Context, dto *domain.DeleteMessageDTO) error {
	msg := domain.Message{ID: dto.ID}
	return s.db.WithContext(ctx).Delete(&msg).Error
}

func (s *messageRepo) Update(ctx context.Context, dto *domain.UpdateMessageDTO) error {
	return s.db.WithContext(ctx).Model(dto.OldMessage).Updates(dto.NewMessage).Error
}

func (s *messageRepo) FindByID(ctx context.Context, id uint) (*domain.Message, error) {
	msg := domain.Message{}
	err := s.db.WithContext(ctx).Take(&msg, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repository.ErrUnknownMessage
		}
		return nil, repository.UnknownErr(err)
	}
	return &msg, nil
}

func (s *messageRepo) FindByChat(ctx context.Context, dto *domain.FindMessageByChatDTO) ([]domain.Message, error) {
	var msgs []domain.Message
	err := s.db.WithContext(ctx).Offset(dto.Offset).Limit(dto.Limit).Find(&msgs).Error
	if err != nil {
		return nil, repository.UnknownErr(err)
	}
	return msgs, nil
}

func NewMessageRepository(gormDB *gorm.DB, logger *logging.Logger) repository.MessageRepository {
	return &messageRepo{
		db:     gormDB,
		logger: logger,
	}
}
