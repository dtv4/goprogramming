package shape

import "fmt"

const (
	CIRCLE    = "Circle"
	RECTANGLE = "Rectangle"
)

type Shape interface {
	draw()
}

type Circle struct{}

func (c Circle) draw() {
	fmt.Println("draw", CIRCLE)
}

type Rectangle struct{}

func (r Rectangle) draw() {
	fmt.Println("draw", RECTANGLE)
}

type ShapeFactory struct{}

func (f ShapeFactory) GetShape(s string) Shape {
	switch s {
	case CIRCLE:
		return Circle{}
	case RECTANGLE:
		return Rectangle{}
	default:
		return nil
	}
}

func FactoryShape() {
	factory := ShapeFactory{}
	circle := factory.GetShape(CIRCLE)
	circle.draw()

	rect := factory.GetShape(RECTANGLE)
	rect.draw()
}
