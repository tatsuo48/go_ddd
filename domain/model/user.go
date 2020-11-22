package model

import (
	"errors"

	"github.com/google/uuid"
)

type UserID string
type UserName string

type User struct {
	userID   UserID
	userName UserName
}

func NewUser(name string) (*User, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return &User{}, errors.New("UUIDの生成に失敗しました。")
	}
	return &User{
		userID:   UserID(u.String()),
		userName: UserName(name),
	}, nil
}

type IUserRepositry interface {
	Save(User) error
	Find(UserID) (*User, error)
	FindByName(UserName) *User
}
