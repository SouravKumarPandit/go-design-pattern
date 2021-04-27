package main

import "strings"

type CustomerInterface interface {
	isNil() bool
	getName() string
}
type AbstractCustomer struct {
	name string
	CustomerInterface
}

type RealCustomer struct {
	AbstractCustomer
}

func NewRealCustomer(name string) RealCustomer {
	customer := RealCustomer{AbstractCustomer{name: name}}
	return customer

}

func (rc RealCustomer) isNil() bool {
	return false
}
func (rc RealCustomer) getName() string {
	return rc.name
}


type NullCustomer struct {
	AbstractCustomer
}

func NewNullCustomer() NullCustomer {
	customer := NullCustomer{AbstractCustomer{}}
	return customer

}

func (rc NullCustomer) isNil() bool {
	return true
}
func (rc NullCustomer) getName() string {
	return "Not Available in Customer Database";
}

type CustomerFactory struct {
	names []string
	
}

func (cf CustomerFactory) getCustomer(name string) CustomerInterface {
	for _, customerName := range cf.names {
		if strings.EqualFold(customerName,name) {
			return NewRealCustomer(name)
		}
	}
	return  NewNullCustomer()
}

func NewCustomerFactory() CustomerFactory {
	return CustomerFactory{names:[]string{"Rob", "Joe", "Julie"} }
}