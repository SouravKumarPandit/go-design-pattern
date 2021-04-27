package main

import "fmt"

type DrawApi interface {
	drawCircle(radius, x, y int)
}
type Shape interface {
	draw()
}
type RedCircle struct {
	radius int
	x      int
	y      int
}

func (rc RedCircle) draw() {
	rc.drawCircle(rc.radius,rc.x,rc.y)
}
func (rc *RedCircle) drawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle [ color : red ] radius %v  point x :%v point y :%v\n" ,radius,x,y)
}


type GreenCircle struct {
	radius int
	x      int
	y      int
}

func (rc GreenCircle) draw() {
	rc.drawCircle(rc.radius,rc.x,rc.y)
}
func (rc *GreenCircle) drawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle [ color : green ] radius %v  point x :%v point y :%v\n" ,radius,x,y)
}

func NewRedCircle(radius, x, y int) Shape {
	return RedCircle{radius, x, y }
}

func NewGreenCircle(radius, x, y int) Shape {
	return GreenCircle{radius, x, y}
}
