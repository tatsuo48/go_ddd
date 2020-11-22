package main

import (
	"fmt"
)

func main() {
	a, _ := NewCircleName("test")
	b, _ := NewCircleName("test")
	fmt.Println(a.Equals(b))
	u, _ := NewUser("test")
	fmt.Println(u)
}
