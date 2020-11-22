package model

type IUserRepositry interface {
	Save(*User)
}
