package main

func main() {
	service := NewServiceLocator().getService("Service1")
	service.execute()
	service = NewServiceLocator().getService("Service2")
	service.execute()
	service = NewServiceLocator().getService("Service1")
	service.execute()
	service = NewServiceLocator().getService("Service2")
	service.execute()

}
