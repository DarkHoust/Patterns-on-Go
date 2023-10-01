package main

import "fmt"

type ITaxi interface {
	getName() string
	getPrice() int
}

type Taxi struct {
	name  string
	price int
}

func (t *Taxi) getName() string {
	return t.name
}

func (t *Taxi) getPrice() int {
	return t.price
}

type Sedan struct {
	ITaxi
}

func NewSedan() ITaxi {
	return &Taxi{name: "Kia K5", price: 35000}
}

type Minivan struct {
	ITaxi
}

func NewMinivan() ITaxi {
	return &Taxi{name: "Kia Stellar", price: 50000}
}

func CallTaxi(carType string) (ITaxi, error) {
	if carType == "Sedan" {
		return NewSedan(), nil
	}
	if carType == "Minivan" {
		return NewMinivan(), nil
	}
	return nil, fmt.Errorf("Unavailable option of taxi")
}

func PrintDetail(it ITaxi) {
	fmt.Println("Taxi's car model:", it.getName())
	fmt.Println("Taxi's car price:", it.getPrice())
}

func main() {
	sedan, _ := CallTaxi("Sedan")
	PrintDetail(sedan)

	minivan, _ := CallTaxi("Minivan")
	PrintDetail(minivan)
}
