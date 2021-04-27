package main

import "fmt"

type Shape interface {
	draw()
}

type Rectangle struct {
}
type Square struct {
}
type Circle struct {
}

func (r Rectangle) draw() {
	fmt.Println("Rectangle : Draw()")
}

func (r Square) draw() {
	fmt.Println("Square : Draw()")
}

func (r Circle) draw() {
	fmt.Println("Circle : Draw()")
}

type ShapeMaker struct {
	circle    Circle
	rectangle Rectangle
	square    Square
}

func NewShapeMaker() ShapeMaker {
	return ShapeMaker{circle: Circle{},
		rectangle: Rectangle{},
		square:    Square{}}
}
func (sm ShapeMaker) drawCircle() {
	sm.circle.draw()
}

func (sm ShapeMaker) drawSquare() {
	sm.square.draw()
}
func (sm ShapeMaker) drawRectangle() {
	sm.rectangle.draw()
}

func main() {
	maker := NewShapeMaker()
	maker.drawCircle()
	maker.drawRectangle()
	maker.drawSquare()
}