package car

type Car struct {
	link  string
	title string
	price float64
}

func NewCar(
	link string,
	title string,
	price float64,
) *Car {
	return &Car{
		link:  link,
		title: title,
		price: price,
	}
}
