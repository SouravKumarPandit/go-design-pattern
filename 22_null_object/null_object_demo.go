package main

import "fmt"

func main() {

	customer1:=NewCustomerFactory().getCustomer("Rob")
	customer2:=NewCustomerFactory().getCustomer("Bob")
	customer3:=NewCustomerFactory().getCustomer("Julie")
	customer4:=NewCustomerFactory().getCustomer("Laura")

	fmt.Println("Customer")
	fmt.Println(customer1.getName())
	fmt.Println(customer2.getName())
	fmt.Println(customer3.getName())
	fmt.Println(customer4.getName())

}
