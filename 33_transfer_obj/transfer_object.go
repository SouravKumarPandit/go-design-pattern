package main

import "fmt"

type StudentVO struct {
	name string
	rollNo int
}

func NewStudentVO(name string ,rollNo int) *StudentVO {
	return &StudentVO{
		name: name,
		rollNo: rollNo,
	}
}

type StudentBO  struct {
	students []StudentVO
}

func NewStudentBO () *StudentBO  {
	s1:=StudentVO{"Robert",0}
	s2:=StudentVO{"John",1}
	return &StudentBO {students: []StudentVO{s1,s2}}
}

func (sbo *StudentBO) deleteStudent(student StudentVO) {

	for i, svo := range sbo.students {
		if svo.rollNo==student.rollNo {
			sbo.students=append(sbo.students[:i], sbo.students[i+1:]...)
		}
	}
		fmt.Println("Student : Roll No ",student.rollNo,"Deleted From database")
}


func (sbo *StudentBO) getAllStudents() []StudentVO {
	return sbo.students
}
func (sbo *StudentBO) getStudent(rollNo int) StudentVO {
	var student StudentVO
	for _, item := range sbo.students {
		if  item.rollNo ==rollNo{
			student=item
		}
	}
	return student
}
func (sbo *StudentBO) updateStudent(student StudentVO) StudentVO {
	//var student StudentVO
	for i, item := range sbo.students {
		if  item.rollNo ==student.rollNo{
			sbo.students[i]=student
		}
	}
	fmt.Println("Student: Roll No " , student.rollNo ,", updated in the database")
	return student
}