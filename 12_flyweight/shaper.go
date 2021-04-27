package main

import "fmt"

type Shape interface {
	draw()
}

//const name =

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func NewCircle(x, y, radius int, color string) Circle {
	return Circle{x: x, y: y, radius: radius, color: color}
}
func NewColorCircle(color string) Circle {
	return Circle{color: color}
}

func (c Circle) draw() {
	fmt.Printf("Circle : Draw %v , x:%v ,Y: %v , Radius : %v \n", c.color, c.x, c.y, c.radius)
}

type ShapeFactory struct {
	circleMap map[string]Circle
}

func NewShapeFactory() ShapeFactory {
	return ShapeFactory{circleMap: map[string]Circle{}}
}

func (sf ShapeFactory) getCircle(color string) Circle {
	circle, ok := sf.circleMap[color]
	if !ok {
		circle = NewColorCircle(color)
		sf.circleMap[color] = Circle{}
		fmt.Println("Created Circle for Color : " + color)
	}
	return circle
}
