package main

func main() {
	delegate := NewBusinessDelegate()
	delegate.setServiceType("EJB")
	client := NewClient(&delegate)
	client.doTask()
	delegate.setServiceType("JMS")
	client.doTask()
}
