package model

import (
	"errors"
	"unicode/utf8"

	"github.com/google/uuid"
)

type UserName string
type UserAddress string
type UUID string

type User struct {
	name    UserName
	address UserAddress
	uuid    UUID
}

func NewUser(name UserName, address UserAddress) (*User, error) {
	if utf8.RuneCountInString(string(name)) <= 3 {
		return &User{}, errors.New("ユーザ名は3文字以上です。")
	}
	u, err := uuid.NewRandom()
	if err != nil {
		return &User{}, errors.New("UUIDの生成に失敗しました。")
	}
	uuid := UUID(u.String())
	return &User{
		name:    name,
		address: address,
		uuid:    uuid,
	}, nil
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) Address() UserAddress {
	return u.address
}

func (u *User) UUID() UUID {
	return u.uuid
}
