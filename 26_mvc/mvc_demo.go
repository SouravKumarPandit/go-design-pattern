package main

func main() {
	student := getStudent()
	view := getStudentView()
	controller := NewStudentController(student, view)
	controller.updateView()

	controller.model.name="Jack Spicer"
	controller.updateView()
}

func getStudentView() StudentView {
	return StudentView{}
}

func getStudent() Student {
	return Student{
		name: "Sourav Pandit",
		rollNo: "3212551211",
	}
}