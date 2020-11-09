package model

import (
	"errors"

	"github.com/google/uuid"
)

type IUserFactory interface {
	Create(UserName, UserAddress) (*User, error)
}

type UserFactory struct {
}

func NewUserFactory() IUserFactory {
	return UserFactory{}
}
func (uf UserFactory) Create(name UserName, address UserAddress) (*User, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return &User{}, errors.New("UUIDの生成に失敗しました。")
	}
	uuid := UUID(u.String())
	user, err := NewUser(name, address, uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
