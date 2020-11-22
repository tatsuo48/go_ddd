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
	Find(CircleID) (*Circle, error)
	FindByName(CircleName) *Circle
}

type ICircleFactory interface {
	Create(CircleName, *User) Circle
}

type CircleService struct {
	circleRepositry ICircleRepositry
}

func NewCircleService(cr ICircleRepositry) CircleService {
	return CircleService{
		circleRepositry: cr,
	}
}

func (cs CircleService) Exists(c Circle) bool {
	circle := cs.circleRepositry.FindByName(c.circleName)
	return circle != nil
}

type CircleCreateCommand struct {
	userID     string
	circleName string
}

func NewCircleCreateCommand(id string, name string) CircleCreateCommand {
	return CircleCreateCommand{
		userID:     id,
		circleName: name,
	}
}

type CircleJoinCommand struct {
	userID   string
	circleID string
}

func NewCircleJoinCommand(userID string, circleID string) CircleJoinCommand {
	return CircleJoinCommand{
		userID:   userID,
		circleID: circleID,
	}
}

type CircleApplicationService struct {
	circleRepositry ICircleRepositry
	circleFactory   ICircleFactory
	circleService   CircleService
	userRepositry   IUserRepositry
}

func NewCircleApplicationService(cr ICircleRepositry, cf ICircleFactory, cs CircleService, us IUserRepositry) CircleApplicationService {
	return CircleApplicationService{
		circleRepositry: cr,
		circleFactory:   cf,
		circleService:   cs,
		userRepositry:   us,
	}
}

func (c CircleApplicationService) Create(cmd CircleCreateCommand) error {
	userID := UserID(cmd.userID)
	owner, err := c.userRepositry.Find(userID)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("サークルのオーナーとなるユーザが見つかりませんでした")
	}
	circleName := CircleName(cmd.circleName)
	circle := c.circleFactory.Create(circleName, owner)
	if c.circleService.Exists(circle) {
		return errors.New("すでに同名のサークルが存在します")
	}
	c.circleRepositry.Save(circle)
	return nil
}

func (c CircleApplicationService) Join(cmd CircleJoinCommand) error {
	userID := UserID(cmd.userID)
	user, err := c.userRepositry.Find(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("ユーザが見つかりませんでした")
	}
	circleID := CircleID(cmd.circleID)
	circle, err := c.circleRepositry.Find(circleID)
	if err != nil {
		return err
	}
	if circle == nil {
		return errors.New("サークルが見つかりませんでした")
	}
	if len(circle.members) >= 29 {
		return errors.New("サークルの参加上限人数(30人)に達しています")
	}
	c.circleRepositry.Save(*circle)
	return nil
}
