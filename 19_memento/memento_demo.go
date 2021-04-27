package main

import "fmt"

func main() {
	org := NewOriginator()
	ctr := NewCareTaker()

	org.setState("State #1 ")
	org.setState("State #2 ")
	ctr.add(org.saveStateToMemento())

	org.setState("State #3 ")
	ctr.add(org.saveStateToMemento())

	org.setState("State #4 ")
	fmt.Printf("Current State : %v \n", org.getState())

	org.getStateFromMemento(ctr.get(0))

	fmt.Printf("First Saved State : %v \n", org.getState())

	org.getStateFromMemento(ctr.get(1))

	fmt.Printf("Second saved State : %v \n", org.getState())
}
