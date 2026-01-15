package car

import (
	"fmt"
	"mobilebg2/define"
)

type Car struct {
	Link       string
	Title      string
	Price      float64
	EngineType string
	Horsepower int64
}

func NewCar(
	link string,
	title string,
	price float64,
	engineType string,
	horsepower int64,
) *Car {
	return &Car{
		Link:       link,
		Title:      title,
		Price:      price,
		EngineType: engineType,
		Horsepower: horsepower,
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
		self.Link,
		self.Title,
		self.Price, define.CURRENCY,
		self.EngineType,
		self.Horsepower,
	)
}
