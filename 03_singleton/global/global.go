package global

import "fmt"

type globalSingleObject struct {
	message string
}

var  singleObject  *globalSingleObject=nil

func (g globalSingleObject) ShowMessage() {
	fmt.Println(g.message)
}
func (g *globalSingleObject) ChangeMessage(message string) {
	g.message=message
}

//it's not thread or gorutine safe
func GetInstance() *globalSingleObject {
	if singleObject==nil {
		singleObject = new(globalSingleObject)
		singleObject.message="Some Message"
	}

	return singleObject
}
