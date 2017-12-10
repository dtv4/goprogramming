package car

import "fmt"

/* translate from java code */
/* Object */
type Car struct {
	speed  Speed
	color  Color
	wheels Wheels
}

func (c *Car) Initialize(b CarBuilder) Car {
	c.speed = b.speed
	c.color = b.Color
	c.wheels = b.wheels
	return *c
}
func (c Car) Drive() error {
	fmt.Println(c, "Drive")
	return nil
}
func (c Car) Stop() error {
	fmt.Println(c, "Stop")
	return nil
}

/* Builder */
type CarBuilder struct {
	speed  Speed
	Color  Color
	wheels Wheels
}

func (b *CarBuilder) TopSpeed(s Speed) *CarBuilder {
	b.speed = s
	return b
}
func (b *CarBuilder) SetColor(color Color) *CarBuilder {
	b.Color = color
	return b
}
func (b *CarBuilder) SetWheels(w Wheels) *CarBuilder {
	b.wheels = w
	return b
}
func (b *CarBuilder) Build() Car {
	var c Car
	return c.Initialize(*b)
}

func BuildInJavaStyle() {
	carBuilder := CarBuilder{Color: RedColor}
	car1 := carBuilder.TopSpeed(MPH).SetWheels(SportsWheels).Build()
	car1.Drive()
	car2 := carBuilder.TopSpeed(KPH).SetWheels(SteelWheels).Build()
	car2.Drive()
}
