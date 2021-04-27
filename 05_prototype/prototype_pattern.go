package main

import (
	"fmt"
)

//FAILED GO lang doe's not support inheritance
//---------------------shape
type Shape struct {
	Id        string
	ShapeType string
}

func (s Shape) Draw() {

}

func (s Shape) getType() string {
	return s.ShapeType
}
func (s Shape) getId() string {
	return s.Id
}

func (s *Shape) setId(Id string) {
	s.Id = Id
}
func (s Shape) clone() interface{} {
	var clone interface{}
	clone = s

	return clone
}

//shape--------------

// Rectangle Shape
type Rectangle struct {
	Shape
}

func (s Rectangle) Draw() {
	fmt.Println("Rectangle Draw Method")
}

// Square Shape
type Square struct {
	Shape
}

func NewSquare() interface{} {
	return Square{Shape{ShapeType: "Square"}}
}
func (s Square) Draw() {
	fmt.Println("Square Draw Method")
}

func (c shapeCache) getShape(s string) Shape {

	return Shape{}
}

type shapeCache struct{}

func (c shapeCache) loadCache() {

}

func NewShapeCache() shapeCache {
	return shapeCache{}
}

func main() {
	cache := NewShapeCache()
	cache.loadCache()

	shape := cache.getShape("1")
	fmt.Println("Shape : " + shape.getType())

}
