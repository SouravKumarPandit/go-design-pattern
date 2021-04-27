package main

func main() {
	computer := NewComputer()
	computer.accept(ComputerPartDisplayVisitor{})
}