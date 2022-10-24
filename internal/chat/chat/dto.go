package chat

type CreateChatDTO struct {
	Name string
}

type DeleteChatDTO struct {
	ID uint
}

type GetMessagesDTO struct {
	ID     uint
	Limit  int
	Offset int
}
