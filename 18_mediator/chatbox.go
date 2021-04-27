package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

func (user User) showMessage(message string)  {
	showMessage(user,message)
}
func showMessage(user User,message string)  {
  fmt.Printf("Chat Box Time %v , Message : %v\n",time.Now(),message)
}