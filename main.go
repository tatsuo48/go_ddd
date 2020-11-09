package main

import (
	"fmt"
	"os"

	"github.com/tatsuo48/go_ddd/domain/model"
)

func main() {
	name := model.UserName("taro")
	address := model.UserAddress("example.com")
	user, err := model.NewUser(name, address)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	uas := initUser("test.txt")
	uas.Register(user)
}
