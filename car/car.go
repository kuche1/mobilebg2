package car

import (
	"fmt"
	"mobilebg2/define"
)

type Car struct {
	link       string
	title      string
	price      float64
	engineType string
	horsepower int64
}

func NewCar(
	link string,
	title string,
	price float64,
	engineType string,
	horsepower int64,
) *Car {
	return &Car{
		link:       link,
		title:      title,
		price:      price,
		engineType: engineType,
		horsepower: horsepower,
	}
}

func (self *Car) Sprintf() string {
	return fmt.Sprintf(
		`Car:
    link: %v
    title: %v
    price: %v %v
    engine type: %v
    horsepower: %v
`,
		self.link,
		self.title,
		self.price, define.CURRENCY,
		self.engineType,
		self.horsepower,
	)
}
