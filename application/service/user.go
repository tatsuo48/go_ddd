package service

import "github.com/tatsuo48/go_ddd/domain/model"

type UserApplicationService struct {
	UserRepositry model.IUserRepositry
}

func NewUserApplicationService(repo model.IUserRepositry) *UserApplicationService {
	return &UserApplicationService{
		UserRepositry: repo,
	}
}

func (uas *UserApplicationService) Register(user *model.User) {
	uas.UserRepositry.Save(user)
}
