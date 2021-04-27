package main

type Student struct {
	name   string
	rollNo int
}

func NewStudent(name string, rollNo int) Student {
	return Student{}
}

type StudentDao interface {
	getAllStudents() []Student
	getStudent(rollNo int) Student
	updateStudent(student Student)
	deleteStudent(student Student)
}

type StudentDaoImpl struct {
	students []Student
}

func NewStudentDaoImpl() *StudentDaoImpl {
	daoImpl := StudentDaoImpl{}
	daoImpl.students = append(daoImpl.students, Student{
		name:   "Robert",
		rollNo: 0,
	})
	daoImpl.students = append(daoImpl.students, Student{
		name:   "John",
		rollNo: 1,
	})
	return &daoImpl
}

func (imp *StudentDaoImpl) deleteStudent(student Student) {
	for i, data := range imp.students {
		if data.name ==student.name && data.rollNo==student.rollNo {
			// Remove the element at index i from a.
			copy(imp.students[i:], imp.students[i+1:]) // Shift a[i+1:] left one index.
			imp.students[len(imp.students)-1] = Student{}    // Erase last element (write zero value).
			imp.students = imp.students[:len(imp.students)-1]
		}
	}
}
func (imp *StudentDaoImpl) updateStudent(student Student) {
	for i, data := range imp.students {
		if data.rollNo==student.rollNo {
			imp.students[i]=student
		}
	}
}
func (imp *StudentDaoImpl) getStudent(rollNo int) Student {
	var student Student
	for _, data := range imp.students{
		if data.rollNo==rollNo{
			student =data
		}


	}
	return student

}
func (imp *StudentDaoImpl) getAllStudents() []Student {

	return imp.students
}
