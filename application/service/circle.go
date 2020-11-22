package service

import (
	"errors"

	"github.com/tatsuo48/go_ddd/domain/model"
	"github.com/tatsuo48/go_ddd/domain/service"
)

type CircleApplicationService struct {
	circleRepositry model.ICircleRepositry
	circleFactory   model.ICircleFactory
	circleService   service.CircleService
	userRepositry   model.IUserRepositry
}

func NewCircleApplicationService(cr model.ICircleRepositry, cf model.ICircleFactory, cs service.CircleService, us model.IUserRepositry) CircleApplicationService {
	return CircleApplicationService{
		circleRepositry: cr,
		circleFactory:   cf,
		circleService:   cs,
		userRepositry:   us,
	}
}

func (c CircleApplicationService) Create(cmd CircleCreateCommand) error {
	userID := model.UserID(cmd.userID)
	owner, err := c.userRepositry.Find(userID)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("サークルのオーナーとなるユーザが見つかりませんでした")
	}
	circleName := model.CircleName(cmd.circleName)
	circle := c.circleFactory.Create(circleName, owner)
	if c.circleService.Exists(circle) {
		return errors.New("すでに同名のサークルが存在します")
	}
	c.circleRepositry.Save(circle)
	return nil
}

func (c CircleApplicationService) Join(cmd CircleJoinCommand) error {
	userID := model.UserID(cmd.userID)
	user, err := c.userRepositry.Find(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("ユーザが見つかりませんでした")
	}
	circleID := model.CircleID(cmd.circleID)
	circle, err := c.circleRepositry.Find(circleID)
	if err != nil {
		return err
	}
	if circle == nil {
		return errors.New("サークルが見つかりませんでした")
	}
	if len(circle.Members) >= 29 {
		return errors.New("サークルの参加上限人数(30人)に達しています")
	}
	c.circleRepositry.Save(*circle)
	return nil
}
