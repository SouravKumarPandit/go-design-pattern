package main

import "fmt"

func main() {

	//bug but can understand concept
	subject :=Subject{state: 25}

	NewHexaObserver(&subject)
	NewOctalObserver(&subject)
	NewBinaryObserver(&subject)

	fmt.Println("First State Change 15")
	subject.setState(15)
	fmt.Println("Second State Change 20")
	subject.setState(20)

}