package main

import "fmt"

type Shape interface {
	draw()
}
type Rectangle struct {
}
func NewRectangle() Rectangle{
	return Rectangle{}
}

func (r Rectangle) draw() {
	fmt.Println("Shape : Rectangle")
}


type Circle struct {
}

func NewCircle() Circle{
	return Circle{}
}
func (r Circle) draw() {
	fmt.Println("Shape : Circle")
}
type ShapeDecorator struct {
	shape Shape
	decoratorDraw func ()
}

func NewShapeDecorator(shape Shape) ShapeDecorator {
	return ShapeDecorator{shape:shape,decoratorDraw: setRedBorder}

}


func (sd ShapeDecorator) draw() {
	sd.shape.draw()
	sd.decoratorDraw()
}
func setRedBorder()  {
	fmt.Println("Border Color : Red ")

}