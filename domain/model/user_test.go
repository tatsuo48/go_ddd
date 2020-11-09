package model

import (
	"testing"
)

func TestUserCreateSuccessCase(t *testing.T) {
	name := UserName("taro")
	address := UserAddress("example.com")
	_, err := NewUser(name, address)
	if err != nil {
		t.Error("user create failed; expected success")
	}
}

func TestUserCreateFailedCase(t *testing.T) {
	name := UserName("sak")
	address := UserAddress("example.com")
	_, err := NewUser(name, address)
	if err == nil {
		t.Errorf("user create success; expected failed")
	}
}
