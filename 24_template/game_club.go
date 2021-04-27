package main

import "fmt"

type IGame interface {
	initialize()
	startPlay()
	endPlay()
}

type Cricket struct {
}

func (c Cricket) play() {
	c.initialize()
	c.startPlay()
	c.endPlay()
}
func (c Cricket) initialize() {
	fmt.Println("Cricket Game Initialized ! Start Playing.")
}

func (c Cricket) startPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game !")

}
func (c Cricket) endPlay() {
	fmt.Printf("Cricket Game Finished !")
}


type Football  struct {
}

func (c Football ) play() {
	c.initialize()
	c.startPlay()
	c.endPlay()
}
func (c Football ) initialize() {
	fmt.Println("Football Game Initialized ! Start Playing.")
}

func (c Football ) startPlay() {
	fmt.Println("Football Game Started. Enjoy the game !")

}
func (c Football ) endPlay() {
	fmt.Printf("Football Game Finished !")
}