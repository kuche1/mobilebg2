package car

type Car struct {
	link  string
	title string
}

func NewCar(link string, title string) *Car {
	return &Car{
		link,
		title,
	}
}
