package main

import "fmt"

func main() {
	impl := NewStudentDaoImpl()
	for _, student := range impl.students {
		fmt.Printf("Student : [Roll No : %v ] ,[Name : %v]]\n",student.rollNo,student.name)
	}
	student1:=impl.getAllStudents()[0]
	student1.name="Michael"
	impl.updateStudent(student1)

	studentUpdated := impl.getAllStudents()[0]
	fmt.Println(studentUpdated)
}
