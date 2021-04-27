package main

import "fmt"

func main() {
	var persons []Person
	fmt.Println(persons)
	persons = append(persons, NewPerson("sourav 1", "male", "single"))
	persons = append(persons, NewPerson("Sima 1", "female", "married"))
	persons = append(persons, NewPerson("sourav 1", "male", "married"))
	persons = append(persons, NewPerson("sourav Pandit", "male", "single"))
	persons = append(persons, NewPerson("Akasha", "female", "single"))
	persons = append(persons, NewPerson("sourav 1", "male", "single"))

	maleCriteria := NewMaleCriteria()
	femaleCriteria := NewFemaleCriteria()
	singleCriteria := NewSingleCriteria()
	singleAndMaleCriteria := NewAddCriteria(singleCriteria, maleCriteria)
	singleOrFemaleCriteria := NewORCriteria(singleCriteria, femaleCriteria)

	fmt.Println("----------------------MALE-----------------------------")
	printPerson(maleCriteria.meetCriteria(persons))
	fmt.Println("----------------------FEMALE-----------------------------")
	printPerson(femaleCriteria.meetCriteria(persons))
	fmt.Println("----------------------SINGLE-----------------------------")
	printPerson(singleCriteria.meetCriteria(persons))
	fmt.Println("----------------------SINGLE MALE-----------------------------")
	printPerson(singleAndMaleCriteria.meetCriteria(persons))
	fmt.Println("----------------------SINGLE OR FEMALE-----------------------------")
	printPerson(singleOrFemaleCriteria.meetCriteria(persons))

}

func printPerson(persons []Person) {

	for _, person := range persons {
		fmt.Printf("Person Info  Name : %v Gender: %v Status: %v\n", person.name, person.gender, person.maritalStatus)
	}

}
