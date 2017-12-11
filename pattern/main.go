package main

import (
	"fmt"
	"my_github/goprogramming/pattern/behavior/number_base"
	"my_github/goprogramming/pattern/builder/car"
	"my_github/goprogramming/pattern/builder/meal"
	"my_github/goprogramming/pattern/factory/shape"
)

func main() {
	fmt.Println("-Builder")
	car.BuildInJavaStyle()
	meal.BuildMeal()
	fmt.Println("-Factory")
	shape.FactoryShape()
	fmt.Println("-Observe")
	base.ObserveNumber()
}
