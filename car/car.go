package car

import "fmt"

type Car struct {
	link       string
	title      string
	price      float64
	engineType string
}

func NewCar(
	link string,
	title string,
	price float64,
	engineType string,
) *Car {
	return &Car{
		link:       link,
		title:      title,
		price:      price,
		engineType: engineType,
	}
}

func (self *Car) Sprintf() string {
	return fmt.Sprintf(
		`Car:
    link: %v
    title: %v
    price: %v
    engine type: %v
`,
		self.link,
		self.title,
		self.price,
		self.engineType,
	)
}
