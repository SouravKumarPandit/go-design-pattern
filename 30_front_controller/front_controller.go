package main

import (
	"fmt"
	"strings"
)

type HomeView struct {
	show func()
}

func NewHomeView() HomeView {
	return HomeView{show: func() {
		fmt.Println("Displaying Home Page ")

	}}
}

type StudentView struct {
	show func()
}

func NewStudentView() StudentView {
	return StudentView{show: func() {
		fmt.Println("Displaying Student Page ")

	}}
}

type Dispatcher struct {
	studentView StudentView
	homeView    HomeView
}

func NewDispatcher() Dispatcher {
	return Dispatcher{homeView: NewHomeView(), studentView: NewStudentView()}
}

func (dis Dispatcher) dispatch(request string) {
	if strings.EqualFold("STUDENT", request) {
		dis.studentView.show()
	} else {
		dis.homeView.show()
	}

}

type FrontController struct {
	dispatcher Dispatcher
}

func NewFrontController() FrontController {
	return FrontController{dispatcher: NewDispatcher()}
}
func (f FrontController) isAuthenticUser() bool {
	fmt.Println("\nUser is authenticated Successfully.")
	return true
}
func (f FrontController) trackRequest(request string) {
	fmt.Printf("Page Requsted : %v",request)
}

func (f FrontController) dispatchRequest(request string)  {
	f.trackRequest(request)
	if f.isAuthenticUser() {
		f.dispatcher.dispatch(request)
	}
}