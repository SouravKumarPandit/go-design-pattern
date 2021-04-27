package main

import "fmt"

func main() {
	//Image image = new ProxyImage("test_10mb.jpg");

	//image will be loaded from disk
	//image.display();
	//System.out.println("");

	//image will not be loaded from disk
	//image.display();
	poximg := NewProxyImage("test_100mb.jpg")
	poximg.display()
	fmt.Println()
	poximg.display()

}
