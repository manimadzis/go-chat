package postgres

import (
	"context"
	"go-chat/internal/chat/entities"
	"go-chat/internal/chat/user"
	"go-chat/internal/chat/user/db"
	"go-chat/pkg/gormclient"
	"os"
	"testing"
)

var createUserDTOs = []user.CreateUserDTO{
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

var users = []entities.User{
	{
		ID:    1,
		Login: "asdb",
	},
	{
		ID:    2,
		Login: "fghgj",
	},
	{
		ID:    3,
		Login: "rtyr",
	},
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
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
	for _, createUser := range createUserDTOs {
		_, err := repo.Create(context.Background(), &createUser)
		if err != nil {
			t.Errorf("can't create user: %v", err)
			continue
		}
	}
}

func TestStorage_FindById(t *testing.T) {
	gromDB, err := gormclient.NewMock()
	if err != nil {
		t.Fatalf("can't open connection to gorm gromDB: %v", err)
	}
	repo := NewRepository(gromDB)
	ctx := context.Background()

	tests := []struct {
		user entities.User
		err  error
	}{
		{
			user: users[0],
			err:  nil,
		},
		{
			user: users[1],
			err:  nil,
		},
		{
			user: entities.User{ID: 1000, Login: "empty"},
			err:  db.ErrUnknownUser,
		},
	}

	for _, test := range tests {
		user, err := repo.FindById(ctx, test.user.ID)
		if err != test.err {
			t.Errorf("Error occures while FindById: %v", err)
			continue
		}

		if test.err == nil && (user.Login != test.user.Login) {
			t.Errorf("invalid result: expected: %v; got: %v", test, user)
		}
	}
}
