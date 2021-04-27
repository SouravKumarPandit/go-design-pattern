package main

import "fmt"

type Student struct {
	rollNo string
	name string
}

func (s Student) getRollNo() string {
	return s.rollNo
}
func (s *Student) setRollNo(rollNo string) {
	s.rollNo=rollNo
}
func (s Student) getName() string {
	return s.name
}
func (s *Student) setName(name string) {
	s.name=name
}

type StudentView struct {
}

func (sv StudentView) printStudentDetails(studentName string, studentRollNo string) {
	fmt.Println("\nStudent : ")
	fmt.Printf(" name : %v",studentName)
	fmt.Printf("	roll no : %v",studentRollNo)

}

type StudentController struct {
	model Student
	view StudentView
}

func NewStudentController(model Student, view StudentView) StudentController {
	return StudentController{model: model,view: view}
}

func (sc StudentController) updateView() {
	sc.view.printStudentDetails(sc.model.getName(),sc.model.getRollNo())
}
