package main

import "fmt"

type ComputerPartVisitor struct {
}

func (cpv2 ComputerPartVisitor) accept(cpv ComputerPartVisitor) {
	panic("implement me")
}

type ComputerPart interface {
	accept(cpv ComputerPartVisitor)
}

func (cpv ComputerPartVisitor) visit(comp ComputerPart) {
	comp.accept(cpv)
}

type KeyBoard struct {
}

func (key KeyBoard) accept(cpv ComputerPartVisitor) {
	cpv.visit(cpv)
}

type Mouse struct {
}

func (key Mouse) accept(cpv ComputerPartVisitor) {
	cpv.visit(cpv)
}

type Monitor struct {
}

func (key Monitor) accept(cpv ComputerPartVisitor) {
	cpv.visit(cpv)
}

type Computer struct {
	parts []ComputerPart
}

func NewComputer() Computer {
	return Computer{parts: []ComputerPart{Mouse{}, KeyBoard{}, Monitor{}}}
}

func (c Computer) accept(cpv IComputerPartVisitor)  {
	for i := 0; i < len(c.parts); i++ {
		part := c.parts[i]
		//part.accept(cpv)

		switch part.(type) {
		case KeyBoard: cpv.visitKeyBord(part.(KeyBoard))
		case Monitor: cpv.visitMonitor(part.(Monitor))
		case Mouse: cpv.visitMouse(part.(Mouse))
		//case Computer: cpv.visitComputer(part.(Computer))

		}
	}
	cpv.visitComputer(c)

}

type IComputerPartVisitor interface {
	visitComputer(computer Computer)
	visitMouse(mouse Mouse)
	visitKeyBord(keybord KeyBoard)
	visitMonitor(montior Monitor)
}

type ComputerPartDisplayVisitor struct {
}

func (cdv ComputerPartDisplayVisitor) visitComputer(computer Computer)  {
	fmt.Println("Displaying Computer")

}

func (cdv ComputerPartDisplayVisitor) visitMouse(mouse Mouse){
	fmt.Println("Displaying Mouse")

}

func (cdv ComputerPartDisplayVisitor) visitKeyBord(keybord KeyBoard){
	fmt.Println("Displaying KeyBord")

}

func (cdv ComputerPartDisplayVisitor) visitMonitor(montior Monitor) {
	fmt.Println("Displaying Monitor")
}