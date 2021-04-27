package main

import "fmt"

func main() {
	cxt:=NewSContext(OperationAdd{})
	printOperation(cxt,"+",10,5)

	cxt=NewSContext(OperationSubstract{})
	printOperation(cxt,"-",10,5)

	cxt=NewSContext(OperationMultiply{})
	printOperation(cxt,"*",10,5)
}

func printOperation(cxt SContext, s string, i int, i2 int) {
	fmt.Printf(" %v %v %v :  %v  \n",i,s,i2,cxt.strategy.doOperation(i,i2))
}