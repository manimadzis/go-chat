package domain

type CreateMessageDTO struct {
	Text   string
	UserID uint
	ChatID uint
}

type DeleteMessageDTO struct {
	ID uint
}

type UpdateMessageDTO struct {
	OldMessage *Message
	NewMessage *Message
}

type FindMessageByChatDTO struct {
	Chat   *Chat
	Limit  int
	Offset int
}
