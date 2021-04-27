package main

import (
	"design.pattern.go/03_singleton/global"
	"fmt"
)

func main() {
	instance1 := global.GetInstance()
	instance1.ShowMessage()


	instance2 := global.GetInstance()
	instance2.ShowMessage()

	fmt.Println("-----------------------------------------")

	//instance1.Message="Message Change in Instance 1"

	instance1.ChangeMessage("Message Change in Instance 1")

	instance1.ShowMessage()
	instance2.ShowMessage()



}
