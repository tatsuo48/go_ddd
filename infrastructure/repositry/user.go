package repositry

import (
	"fmt"

	"github.com/tatsuo48/go_ddd/domain/model"
)

type InMemoryUserRepositry struct {
	filename string
}

func NewInMemoryUserRepositry(filename string) model.IUserRepositry {
	return InMemoryUserRepositry{
		filename: filename,
	}
}

func (ur InMemoryUserRepositry) Save(user *model.User) {
	fmt.Printf("save %s, UUID is %s", user.Name(), user.UUID())
}
