package main

import (
	"github.com/tatsuo48/go_ddd/application/command"
)

func main() {
	cmd := command.UserRegisterCommand{
		Name:    "taro",
		Address: "example.com",
	}
	uas := initUserApplicationService("test.txt")
	uas.Register(cmd)
}
