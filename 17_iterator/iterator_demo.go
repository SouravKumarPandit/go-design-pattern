package main

import "fmt"

func main() {
	repo := NewNameRepository([]string{"Robert", "John", "Julie", "Lora"})

	for repo.hasNext() {
		name := repo.next()
		fmt.Printf("Name : %v \n", name)
	}
}
