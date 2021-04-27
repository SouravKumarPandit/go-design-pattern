package globalsync

import (
	"fmt"
	"sync"
)

type globalSingleObject struct {
	message string
}

var  singleObject  *globalSingleObject=nil
var once sync.Once

func (g globalSingleObject) ShowMessage() {
	fmt.Println(g.message)
}
func (g *globalSingleObject) ChangeMessage(message string) {
	g.message=message
}

//it's  thread or gorutine safe
func GetInstance() *globalSingleObject {
	once.Do(func() {
		singleObject = &globalSingleObject{}
	})

	return singleObject
}
