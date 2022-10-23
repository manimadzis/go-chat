package postgres

import (
	"context"
	"fmt"
	"go-chat/internal/chat/user"
	"go-chat/pkg/gormclient"
	"testing"
)

var users = []user.CreateUserDTO{
	{
		Login:    "asdb",
		Password: "123",
	},
	{
		Login:    "fghgj",
		Password: "4568",
	},
	{
		Login:    "rtyr",
		Password: "888",
	},
}

func TestStorage_AutoMigrate(t *testing.T) {
	db, err := gormclient.NewMock()
	if err != nil {
		t.Fatalf("can't open connection to gorm db: %v", err)
	}
	repo := NewRepository(db)

	err = repo.AutoMigrate(context.Background())
	if err != nil {
		t.Errorf("can't migrate: %v", err)
	}
}

func TestStorage_Create(t *testing.T) {
	db, err := gormclient.NewMock()
	if err != nil {
		t.Fatalf("can't open connection to gorm db: %v", err)
	}
	repo := NewRepository(db)
	for _, user := range users {
		_, err := repo.Create(context.Background(), &user)
		if err != nil {
			t.Errorf("can't create user: %v", err)
		}
	}
}

func TestStorage_FindById(t *testing.T) {
	db, err := gormclient.NewMock()
	if err != nil {
		t.Fatalf("can't open connection to gorm db: %v", err)
	}
	repo := NewRepository(db)

	ctx := context.Background()
	u, err := repo.FindById(ctx, 1)
	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Println(u)
}
