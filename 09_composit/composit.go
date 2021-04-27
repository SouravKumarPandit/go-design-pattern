package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	ceo := NewEmployee("sourav Kumar", "CEO", 1000000)
	headSale := NewEmployee("Robert", "Head Sales", 200000)

	headMarketing := NewEmployee("Michel", "Head Marketing", 200000)

	clerk1 := NewEmployee("Laura", "Head Sales", 100000)
	clerk2 := NewEmployee("Bob", "Head Sales", 100000)

	saleExecutive1 := NewEmployee("Richer", "Sales", 120000)
	saleExecutive2 := NewEmployee("Rob", "Sales", 120000)


	headSale.Add(saleExecutive1)
	headSale.Add(saleExecutive2)

	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	ceo.Add(headSale)
	ceo.Add(headMarketing)
	//fmt.Println(ceo)


	var jsonData []byte
	jsonData, err := json.Marshal(&struct {
		name string `json:"name"`
	}{name: "Sourav Kumar Pandit"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))


	/*jsonByte, err := json.Marshal(ceo)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(jsonByte)*/

	/*for _, subordinate := range ceo.GetSubordinates() {
		fmt.Println(subordinate)
		for _, employee := range subordinate.subordinates {
			if subordinate.subordinates != nil {
				fmt.Println(employee)
			}

		}

	}*/

}
