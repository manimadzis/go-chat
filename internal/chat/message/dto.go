package message

import "go-chat/internal/chat/entities"

type CreateMessageDTO struct {
	Text   string
	UserID uint
	ChatID uint
}

type DeleteMessageDTO struct {
	ID uint
}

type UpdateMessageDTO struct {
	OldMessage *entities.Message
	NewMessage *entities.Message
}

type FindByChatDTO struct {
	Chat   *entities.Chat
	Limit  int
	Offset int
}
