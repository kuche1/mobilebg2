package car

import (
	"fmt"
	"mobilebg2/define"
)

type Car struct {
	Link         string
	Title        string
	Price        float64
	EngineType   string
	Horsepower   int64
	YearProduced int16
	Mialage      int64
}

func NewCar(
	link string,
	title string,
	price float64,
	engineType string,
	horsepower int64,
	yearProduced int16,
	mialage int64,
) *Car {
	return &Car{
		Link:         link,
		Title:        title,
		Price:        price,
		EngineType:   engineType,
		Horsepower:   horsepower,
		YearProduced: yearProduced,
		Mialage:      mialage,
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
    year produced: %v
    mialage: %v km
`,
		self.Link,
		self.Title,
		self.Price, define.CURRENCY,
		self.EngineType,
		self.Horsepower,
		self.YearProduced,
		self.Mialage,
	)
}
