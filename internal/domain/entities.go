package domain

import (
	"time"
)

type Table interface {
	TableName() string
}

type User struct {
	ID           uint      `db:"id" json:"id,omitempty" gorm:"primaryKey"`
	Login        string    `db:"login" json:"login,omitempty" gorm:"uniqueIndex;not null"`
	PasswordHash string    `db:"password_hash" json:"-" gorm:"not null"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	ExpiredAt    time.Time `db:"expire_at" json:"expire_at" gorm:"type:timestamp"`
	IsDeleted    bool      `db:"is_deleted" json:"-" gorm:"default:false"`

	Messages []Message `json:"-" gorm:"foreignKey:UserID"`
	Chats    []Chat    `json:"-" gorm:"many2many:user_chat"`
}

type Chat struct {
	ID      uint   `db:"id" json:"id,omitempty" gorm:"primaryKey"`
	Name    string `db:"name" json:"name,omitempty" gorm:"not null"`
	OwnerID uint   `db:"owner" gorm:"not null"`

	Owner    *User     `json:"owner" gorm:"foreignKey:OwnerID"`
	Users    []User    `json:"-" gorm:"many2many:user_chat"`
	Messages []Message `json:"-" gorm:"foreignKey:ChatID"`
}

type Message struct {
	ID     uint      `db:"id" json:"id,omitempty" gorm:"primaryKey"`
	Text   string    `db:"text" json:"text,omitempty" gorm:"not null"`
	Time   time.Time `db:"timestamp" json:"timestamp,omitempty" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ChatID uint      `db:"chat_id" json:"chat_id" gorm:"not null"`
	UserID uint      `db:"user_id" json:"author_id,omitempty" gorm:"not null"`

	Chat *Chat `json:"-" gorm:"foreignKey:ChatID"`
	User *User `json:"-" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}

func (Chat) TableName() string {
	return "chat"
}

func (Message) TableName() string {
	return "message"
}
