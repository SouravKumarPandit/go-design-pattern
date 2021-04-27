package main

import "fmt"

func main() {
	isMale := getMaleExpression()
	isMarriedWoman := getMarriedWomanExpression()
	fmt.Printf("John is Male ? %v \n",
		isMale.interpret("Jhon"))

	fmt.Printf("Julie is Married Wommen ? %v",
		isMarriedWoman.interpret("Married Julie"))
}

func getMaleExpression() Expression {
	robert := NewTerminalExpression("Robert")
	john := NewTerminalExpression("Jhon")
	return NewOrExpression(robert, john)
}

func getMarriedWomanExpression() Expression {
	julie := NewTerminalExpression("Julie")
	married := NewTerminalExpression("Married")
	return NewAndExpression(julie, married)
}
