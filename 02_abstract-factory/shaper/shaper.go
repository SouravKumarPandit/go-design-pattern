package shaper

import "fmt"

//shape enum
type ShapeType int

const (
	ROUND_RECTANGLE ShapeType = iota
	ROUND_SQUARE
	RECTANGLE
	SQUARE
)

// validate enum : change validation based on name
func (st ShapeType) IsValid() bool {
	return int(st) >= int(ROUND_RECTANGLE) && int(st) <= int(SQUARE)
}

// shape interface

type Shape interface {
	Draw()
}

type RoundRectangle struct {
}

func (s RoundRectangle) Draw() {
	fmt.Println("RoundRectangle Draw Method")
}

type RoundSquare struct {
}

func (s RoundSquare) Draw() {
	fmt.Println("RoundSquare Draw Method")
}

type Rectangle struct {
}

func (s Rectangle) Draw() {
	fmt.Println("Rectangle Draw Method")
}

type Square struct {
}

func (s Square) Draw() {
	fmt.Println("Square Draw Method")
}

// abstract layout 
type AbstractFactory interface {
	GetShape(st ShapeType) Shape
}
type ShapeFactory struct {
}

func (sf ShapeFactory) GetShape(st ShapeType) Shape {
	if !st.IsValid() {
		//lazy work change validation for rounded only range
		return Rectangle{}
	}

	var s Shape
	switch st {
	case ROUND_RECTANGLE, RECTANGLE:
		s = Rectangle{}
	case ROUND_SQUARE, SQUARE:
		s = Square{}

	}

	return s
}

type RoundShapeFactory struct {
}

func (sf RoundShapeFactory) GetShape(st ShapeType) Shape {
	if !st.IsValid() {
		//lazy work change validation for rounded only range
		return Rectangle{}
	}

	var s Shape
	switch st {
	case ROUND_RECTANGLE, RECTANGLE:
		s = RoundRectangle{}
	case ROUND_SQUARE, SQUARE:
		s = RoundSquare{}
	}

	return s
}
