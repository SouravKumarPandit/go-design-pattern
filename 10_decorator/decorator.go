package main

import "fmt"

func main() {

	shapeCircle:=NewCircle()
	shapeRectangle:=NewRectangle()

	decoratorCircle:=NewShapeDecorator(shapeCircle)
	decoratorRectangle:=NewShapeDecorator(shapeRectangle)

	shapeCircle.draw()

	fmt.Println("----------------------------")
	decoratorCircle.draw()
	fmt.Println("----------------------------")
	decoratorRectangle.draw()

}
