package service

import (
	"github.com/tatsuo48/go_ddd/application/command"
	"github.com/tatsuo48/go_ddd/domain/model"
)

type UserApplicationService struct {
	UserRepositry model.IUserRepositry
	UserFactry    model.IUserFactry
}

func NewUserApplicationService(repo model.IUserRepositry, factry model.IUserFactry) *UserApplicationService {
	return &UserApplicationService{
		UserRepositry: repo,
		UserFactry:    factry,
	}
}

func (uas *UserApplicationService) Register(cmd command.UserRegisterCommand) error {
	name := model.UserName(cmd.Name)
	address := model.UserAddress(cmd.Address)
	user, err := uas.UserFactry.Create(name, address)
	if err != nil {
		return err
	}
	uas.UserRepositry.Save(user)
	return nil
}
