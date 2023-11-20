package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing Circle")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing Rectangle")
}

type AbstractFactory interface {
	CreateCircle() Shape
	CreateRectangle() Shape
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateCircle() Shape {
	return &Circle{}
}

func (f *ConcreteFactory1) CreateRectangle() Shape {
	return &Rectangle{}
}

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateCircle() Shape {
	return &Circle{}
}

func (f *ConcreteFactory2) CreateRectangle() Shape {
	return &Rectangle{}
}

func main() {
	// Using ConcreteFactory1 to create rounded shapes
	factory1 := &ConcreteFactory1{}
	circle1 := factory1.CreateCircle()
	rectangle1 := factory1.CreateRectangle()

	circle1.Draw()
	rectangle1.Draw()

	// Using ConcreteFactory2 to create sharp shapes
	factory2 := &ConcreteFactory2{}
	circle2 := factory2.CreateCircle()
	rectangle2 := factory2.CreateRectangle()

	circle2.Draw()
	rectangle2.Draw()
}
