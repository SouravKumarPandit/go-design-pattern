package order

import (
	"fmt"
)

type Item interface {
	Name() string
	Packing() Packing
	price() float32
}

//----------------packing
type Packing struct {
	Pack string
}
type Wrapper struct{}

func (p Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct{}

func (p Bottle) Pack() string {
	return "Wrapper"
}

//packing----------------

//-------------------item

type Burger struct {
	Wrapper,
	itemPrice float32
}
type ColdDrink struct {
	Wrapper,
	itemPrice float32
}

//item-----------------

type VegBurger struct {
	Burger
}

func (v VegBurger) Name() string {
	return "Veg Burger"
}

func (v VegBurger) Packing() Packing {
	return Packing{Pack: "Wrapper"}

}

func (v VegBurger) price() float32 {
	return 17.12
}

type ChickenBurger struct {
	Burger
}

func (c ChickenBurger) Name() string {
	return "Chicken Burger"
}

func (c ChickenBurger) Packing() Packing {
	return  Packing{Pack: "Wrapper"}
}

func (c ChickenBurger) price() float32 {
	return 25.55
}

type Coke struct {
	ColdDrink
}

func (c Coke) Name() string {
	return "Coke"
}

func (c Coke) Packing() Packing {
	return Packing{Pack: "Bottle"}

}

func (c Coke) price() float32 {
	return 47.00

}

type Pepsi struct {
	ColdDrink
}

func (p Pepsi) Name() string {
	return "Pepsi"
}

func (p Pepsi) Packing() Packing {
	return Packing{Pack: "Bottle"}
}

func (p Pepsi) price() float32 {
	return 45.00
}

//----------------------
type ItemBuilder interface {
	add(item Item)
	cost() float32
	ShowItem() string
}
type Meal struct {
	items []Item
}

func (m *Meal) add(item Item) {
	m.items = append(m.items, item)
	//m.itemBuilder.add=$(m.add)
}

func (m *Meal) cost() float32 {
	return 45.5
}

func (m *Meal) ShowItem() string {
    sb:=string("")
	for _, item := range m.items {
		s:=fmt.Sprintf("Meal Item :  Name : %v \n", item.Name())
		s+=fmt.Sprintf("Meal Item :  Price : %v \n", item.price())
		s+=fmt.Sprintf("Meal Item :  Packing : %v \n", item.Packing().Pack)
		sb+=s
	}

	return sb
}

type MealBuilder interface {
	PrepareVegMeal() Meal
	PrepareNonVegMeal() Meal
}

func NewMealBuilder() MealBuilder {
	var buildMeal = builderPattern{}

	return buildMeal
}

type builderPattern struct{}

func (b builderPattern) PrepareVegMeal() Meal {

	meal := Meal{}
	meal.add(VegBurger{})
	meal.add(Coke{})
	return meal
}

func (b builderPattern) PrepareNonVegMeal() Meal {
	meal := Meal{}
	meal.add(ChickenBurger{})
	meal.add(Pepsi{})
	return meal
}
