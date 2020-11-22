package model

import (
	"testing"
)

func TestUserCreateSuccessCase(t *testing.T) {
	name := UserName("taro")
	address := UserAddress("example.com")
	uuid := UUID("3b81065b-100a-4d23-a67f-38c7ad01e322")
	_, err := NewUser(name, address, uuid)
	if err != nil {
		t.Error("user create failed; expected success")
	}
}

func TestUserCreateFailedCase(t *testing.T) {
	name := UserName("sak")
	address := UserAddress("example.com")
	uuid := UUID("3b81065b-100a-4d23-a67f-38c7ad01e322")
	_, err := NewUser(name, address, uuid)
	if err == nil {
		t.Errorf("user create success; expected failed")
	}
}
