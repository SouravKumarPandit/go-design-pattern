package main

import (
	"fmt"
)

type Observe interface {
	Update()
}

type Observers struct {
	Observe
	subject Subject
}

func (o *Observers) Update() {
	fmt.Printf("Must implement to other\n")
	//o.subject.notifyAllObservers()

}

type Subject struct {
	observers []Observe
	state     int
}

func (s *Subject) getState() int {
	return s.state
}
func (s *Subject) setState(state int) {
	s.state = state
	s.notifyAllObservers()
}
func (s *Subject) attach(observer Observe)  {
	s.observers=append(s.observers, observer)

}

func (s *Subject) notifyAllObservers() {
	for _, observer := range s.observers {
		observer.Update()

		//if reflect.TypeOf(observer)
		//
		//}
	}
}

type BinaryObserver struct {
	Observers
}

func NewBinaryObserver(sub *Subject) BinaryObserver {

	observers := Observers{subject: *sub}

	binaryObserver := BinaryObserver{observers}
	sub.attach(&binaryObserver)
	return binaryObserver
}


func (o *BinaryObserver) Update() {
fmt.Printf("Binary string : %b\n" ,o.subject.getState())
}


type OctalObserver struct {
	Observers
}

func NewOctalObserver(sub *Subject) OctalObserver {

	observers := Observers{subject: *sub}
	observer := OctalObserver{observers}
	sub.attach(&observer)
	return observer
}


func (o *OctalObserver) Update() {
	fmt.Printf("OctalObserver string : %o\n" ,o.subject.getState())
}




type HexaObserver struct {
	Observers
}

func NewHexaObserver(sub *Subject) HexaObserver {
	observers := Observers{subject: *sub}
	observer := HexaObserver{observers}
	sub.attach(&observer)
	return observer
}


func (o *HexaObserver) Update() {
	fmt.Printf("HexaObserver string : %x\n" ,o.subject.getState())
}

