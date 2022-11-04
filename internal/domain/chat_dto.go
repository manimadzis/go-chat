package domain

type DeleteChatDTO struct {
	ID uint
}

type CreateChatDTO struct {
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}

func (c *CreateChatDTO) Valid() error {
	return nil
}

type UpdateChatDTO struct {
	OldChat Chat
	NewChat Chat
}
