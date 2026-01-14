package car

import "fmt"

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

func (self *Car) Sprintf() string {
	return fmt.Sprintf(
		`Car:
    link: %v
    title: %v
    price: %v
`,
		self.link,
		self.title,
		self.price,
	)
}
