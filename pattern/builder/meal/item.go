package meal

type Item interface {
	name() string
	price() float64
}

type Burger struct {
	n string
	p float64
}

func (b Burger) name() string {
	return b.n
}
func (b Burger) price() float64 {
	return b.p
}

var VegBurger = Burger{n: "Veg", p: 10.1}
var ChickenBurger = Burger{n: "Chicken", p: 100.2}

type Drink struct {
	n string
	p float64
}

func (d Drink) name() string {
	return d.n
}
func (d Drink) price() float64 {
	return d.p
}

var Coke = Drink{n: "Coke", p: 11.1}
var Pepsi = Drink{n: "Pepsi", p: 111.2}
