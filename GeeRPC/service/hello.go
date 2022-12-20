package service

import "fmt"

type Hello struct {
}

func (Hello) Say() {
	fmt.Println("hi!")
}
