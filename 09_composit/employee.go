package main

type Employee struct {
	name         string     `json:"name"`
	dept         string     `json:"dept"`
	salary       int        `json:"salary"`
	subordinates []Employee `json:"subordinates"`
}

func NewEmployee(name, dept string, sal int) Employee {
	return Employee{name: name, dept: dept, salary: sal}
}
func (e *Employee) Add(employee Employee) {

	e.subordinates = append(e.subordinates, employee)
}

func (e *Employee) Remove(employee Employee) {
	for i := 0; i < len(e.subordinates); i++ {
		if employee.name == e.name && employee.salary == e.salary && employee.dept == e.dept {
			//e.subordinates = append(e.subordinates[:i], e.subordinates[i+1:])
			copy(e.subordinates[i:], e.subordinates[i+1:])
		}
	}
	//e.subordinates=append(e.subordinates, employee)

}

func (e Employee) GetSubordinates() []Employee {
	return e.subordinates
}

/*
func (e Employee) String() string {
	return fmt.Sprintf("------Employee : %v \t\n %v \t\n %v \t\n %v \n", e.name, e.dept, e.salary, e.subordinates)
}
*/
