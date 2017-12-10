package car

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

/* General Code ? */
type Builder interface {
	TopSpeed(Speed) Builder
	Color(Color) Builder
	Wheels(Wheels) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

func NewBuilder() Builder {
	var b Builder
	return b
}
