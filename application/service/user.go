package service

import (
	"github.com/tatsuo48/go_ddd/application/command"
	"github.com/tatsuo48/go_ddd/domain/model"
)

type UserApplicationService struct {
	UserRepositry model.IUserRepositry
	UserFactory   model.IUserFactory
}

func NewUserApplicationService(repo model.IUserRepositry, Factory model.IUserFactory) *UserApplicationService {
	return &UserApplicationService{
		UserRepositry: repo,
		UserFactory:   Factory,
	}
}

func (uas *UserApplicationService) Register(cmd command.UserRegisterCommand) error {
	name := model.UserName(cmd.Name)
	address := model.UserAddress(cmd.Address)
	user, err := uas.UserFactory.Create(name, address)
	if err != nil {
		return err
	}
	uas.UserRepositry.Save(user)
	return nil
}
