package main

import (
	"fmt"
	"strings"
)

type Service interface {
	getName() string
	execute()
}

type Service1 struct {
}

func (s1 Service1) getName() string {

	return "Service 1 "

}

func (s1 Service1) execute() {
	fmt.Println("Executing Service 1")
}

type Service2 struct {
}

func (s2 Service2) getName() string {

	return "Service 2 "

}

func (s2 Service2) execute() {
	fmt.Println("Executing Service 2")
}

type InitialContext struct {
}

func (ic InitialContext) lookUp(jndiName string) interface{} {
	if strings.EqualFold("SERVICE1", jndiName) {
		fmt.Println("Looking up and creating a new Service1 object")
		return Service1{}
	} else if strings.EqualFold("SERVICE2", jndiName) {
		fmt.Println("Looking up and creating a new Service2 object")
		return Service2{}
	}
	return nil
}

type Cache struct {
	services []Service
}

func NewCache() Cache {
	return Cache{
		services: []Service{},
	}
}

func (cah Cache ) getService(serviceName string) Service {
	for _, service := range cah.services {
		if strings.EqualFold(service.getName(),serviceName) {
			fmt.Println("Returning Cached",serviceName," Object")
			return service
		}

	}
	return nil
}

func (cah *Cache ) addService(newService Service) {
	var exists bool=false
	for _, service := range cah.services {
		if strings.EqualFold(service.getName(),newService.getName()) {
			exists=true
		}
	}
	if !exists {
		cah.services = append(cah.services,newService )
	}
}

type ServiceLocator struct {
	cache Cache
}

func NewServiceLocator() ServiceLocator {
	return  ServiceLocator{cache: Cache{}}
}

func (sl ServiceLocator ) getService(jndiName string) Service {
	service := sl.cache.getService(jndiName)
	if service!=nil {
		return service
	}
	context := InitialContext{}
	up := context.lookUp(jndiName)
	service1 := up.(Service)
	sl.cache.addService(service1)
	return service1
}