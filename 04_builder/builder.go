package main

import (
	order "design.pattern.go/04_builder/item"
	"fmt"
)

func main() {
	f := fmt.Println
	f("Order Meal")

	mb := order.NewMealBuilder()

	meal := mb.PrepareNonVegMeal()
	fmt.Println(meal.ShowItem())

	vegMeal := mb.PrepareVegMeal()
	fmt.Println(vegMeal.ShowItem())

}
