package main

import (
	"github.com/stretchr/stew/slice"
	"strings"
)

type Person struct {
	name          string
	gender        string
	maritalStatus string
}

func NewPerson(name, gender, maritalStatus string) Person {
	return Person{
		name:          name,
		gender:        gender,
		maritalStatus: maritalStatus,
	}
}

type Criteria interface {
	meetCriteria(persons []Person) []Person
}
type CriteriaMale struct {
}

func NewMaleCriteria() Criteria{
	return CriteriaMale{}
}
func (malePerson CriteriaMale) meetCriteria(persons []Person) []Person {
	var filterList []Person
	/*for i := 0; i < len(persons); i++ {
		if persons {

		}
	}*/
	for _, person := range persons {
		if strings.EqualFold(person.gender, "male") {
			filterList = append(filterList, person)
		}

	}
	return filterList
}

type CriteriaFemale struct {
}

func NewFemaleCriteria() Criteria{
	return CriteriaFemale{}
}
func (malePerson CriteriaFemale) meetCriteria(persons []Person) []Person {
	var filterList []Person

	for _, person := range persons {
		if strings.EqualFold(person.gender, "female") {
			filterList = append(filterList, person)
		}

	}
	return filterList
}

type CriteriaSingle struct {
}

func NewSingleCriteria() Criteria{
	return CriteriaSingle{}
}
func (malePerson CriteriaSingle) meetCriteria(persons []Person) []Person {
	var filterList []Person
	for _, person := range persons {
		if strings.EqualFold(person.maritalStatus, "single") {
			filterList = append(filterList, person)
		}

	}
	return filterList
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewAddCriteria(criteria1, criteria2 Criteria) AndCriteria{
	return AndCriteria{criteria: criteria1,otherCriteria: criteria2}
}

func (c AndCriteria) meetCriteria(persons []Person) []Person {
	var filterList =c.criteria.meetCriteria(persons)
	filterAnd := c.otherCriteria.meetCriteria(filterList)

	return filterAnd
}



type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewORCriteria(criteria1, criteria2 Criteria) OrCriteria{
	return OrCriteria{criteria: criteria1,otherCriteria: criteria2}
}

func (c OrCriteria) meetCriteria(persons []Person) []Person {
	var filterList1 =c.criteria.meetCriteria(persons)
	var filterList2 =c.otherCriteria.meetCriteria(persons)
	for _, person := range filterList2 {
		if slice.Contains(filterList1, person) {
			filterList1=append(filterList1, person)
		}

	}
	return filterList1
}
