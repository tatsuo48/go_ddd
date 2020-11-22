package main

import (
	"errors"
	"unicode/utf8"
)

type CircleID string

type CircleName string

func NewCircleName(s string) (CircleName, error) {
	if utf8.RuneCountInString(s) < 3 {
		return CircleName(""), errors.New("サークル名は3文字以上です。")
	} else if utf8.RuneCountInString(s) > 20 {
		return CircleName(""), errors.New("サークル名は20文字以下です。")
	}
	return CircleName(s), nil
}

func (c CircleName) Equals(otherCircleName CircleName) bool {
	return string(c) == string(otherCircleName)
}

type Circle struct {
	circeID    CircleID
	circleName CircleName
	owner      User
	members    []User
}

type ICircleRepositry interface {
	Save(Circle) error
	Find(CircleID) (Circle, error)
	FindByName(CircleName) (Circle, error)
}

type ICircleFactory interface {
	Create(CircleName, User) (Circle, error)
}
