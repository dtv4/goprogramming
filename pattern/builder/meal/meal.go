package meal

import "fmt"

type Meal struct {
	items []Item
}

func (m *Meal) addItem(i Item) {
	m.items = append(m.items, i)
}
func (m *Meal) getCost() float64 {
	var total float64
	for _, item := range m.items {
		total += item.price()
	}
	return total
}
func (m *Meal) showItems() {
	for _, item := range m.items {
		fmt.Println(item.name(), item.price())
	}
}

type MealBuilder struct {
	items []Item
}

func (b *MealBuilder) prepareVegMeal() Meal {
	var m Meal
	b.items = []Item{VegBurger, Coke}
	for _, item := range b.items {
		m.addItem(item)
	}
	return m
}

func BuildMeal() {
	builder := MealBuilder{}
	vegMeal := builder.prepareVegMeal()
	vegMeal.showItems()
	fmt.Println("Total: ", vegMeal.getCost())
}
