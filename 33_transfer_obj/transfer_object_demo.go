package main

import "fmt"

func main() {
	studentObj:=NewStudentBO()
	for _, student := range studentObj.students {

		fmt.Println("Student : [ Roll No : ",student.rollNo ,", Name : ",student.name)

	}

	studneVo:=studentObj.getAllStudents()[0]
	studneVo.name="Michael"

	studentObj.updateStudent(studneVo)
	studneVo=studentObj.getStudent(0)
	fmt.Println("Student : [ Roll No : ",studneVo.rollNo ,", Name : ",studneVo.name)

}