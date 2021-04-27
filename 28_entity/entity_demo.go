package main

func main() {

	/*pending */
	client:=NewClient()
	client.setData("Test","Data")
	client.printData()
	client.setData("Second Test","Data 1")
	client.printData()

}
