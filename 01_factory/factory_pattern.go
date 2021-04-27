package main

import (
	"fmt"
)

//_go:generate enumer -type=ShapeType
type ShapeType int

const (
	RECTANGLE ShapeType = iota
	SQUARE
	CIRCLE
)

func (st ShapeType) IsValid() bool {
	return int(st) >= int(RECTANGLE) && int(st) <= int(CIRCLE)
}

type Shape interface {
	draw()
}

type Rectangle struct{}

func (r Rectangle) draw() {
	fmt.Println("Rectangle Draw Method.")
}

type Square struct{}

func (r Square) draw() {
	fmt.Println("Square Draw Method.")
}

type Circle struct{}

func (r Circle) draw() {
	fmt.Println("Circle Draw Method.")
}
func ShapeFactory(shapeType ShapeType) Shape {
	if !shapeType.IsValid() {
		//you can panic here
		return Rectangle{}
	}

	var s Shape
	switch shapeType {
	case RECTANGLE:
		s = Rectangle{}
		return s
	case SQUARE:
		s = Square{}
		break
	case CIRCLE:
		s = Circle{}
		break
	//default:
	//	s = Rectangle{}
	}

	return s
}

func main() {
	//mf := new(Shape)
	//fmt.Println(mf)

	s := ShapeFactory(RECTANGLE)
	s.draw()

	s = ShapeFactory(SQUARE)
	s.draw()

	s = ShapeFactory(CIRCLE)
	s.draw()
	//can trow null pointer  through error
	s = ShapeFactory(-1)
	s.draw()
}
