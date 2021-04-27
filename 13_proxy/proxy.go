package main

import "fmt"

type Image interface {
	display()
}

type RealImage struct {
	fileName string
}

func NewRealImage(fileName string) RealImage {
	return RealImage{fileName: fileName}
}

func (rm RealImage) display() {
	fmt.Println("Displaying Image : " + rm.fileName)
}

func (rm RealImage) loadFromDisk(fileName string) {
	fmt.Println("Loading : " + fileName)
}

type ProxyImage struct {
	fileName  string
	realImage *RealImage
}

func NewProxyImage(fileName string) ProxyImage {
	return ProxyImage{fileName: fileName}
}

func (rm *ProxyImage) display() {

	if rm.realImage == nil {
		image := NewRealImage(rm.fileName)
		rm.realImage = &image
		rm.realImage.loadFromDisk(rm.fileName)

	}
	fmt.Println("Displaying Image : " + rm.fileName)
}
