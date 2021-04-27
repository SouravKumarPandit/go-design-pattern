package main

func main() {
	controller := NewFrontController()
	controller.dispatchRequest("HOME")
	controller.dispatchRequest("STUDENT")
}
