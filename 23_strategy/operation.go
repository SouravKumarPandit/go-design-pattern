package main

type Strategy interface {
	doOperation(num1, num2 int) int
}

type OperationAdd struct {
}

func (op OperationAdd) doOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationSubstract struct {
}

func (op OperationSubstract) doOperation(num1, num2 int) int {
	return num1 - num2
}

type OperationMultiply struct {
}

func (op OperationMultiply) doOperation(num1, num2 int) int {
	return num1 * num2
}

type SContext struct {
	strategy Strategy
}

func NewSContext(strategy Strategy) SContext {
	return SContext{strategy: strategy}

}





