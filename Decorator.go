package main

import "fmt"

type IBackpack interface {
	getWeight() int
}

type Backpack struct {
}

func (b *Backpack) getWeight() int {
	return 1500
}

type Laptop struct {
	ib IBackpack
}

func (l *Laptop) getWeight() int {
	totalWeight := l.ib.getWeight()
	return totalWeight + 2000
}

type Tablet struct {
	ib IBackpack
}

func (t *Tablet) getWeight() int {
	totalWeight := t.ib.getWeight()
	return totalWeight + 800
}

type BottleOfWater struct {
	ib IBackpack
}

func (bow *BottleOfWater) getWeight() int {
	totalWeight := bow.ib.getWeight()
	return totalWeight + 500
}

func main() {
	backpack := &Backpack{}

	backpackWithLaptor := &Laptop{backpack}
	backpackWithLaptorAndTablet := &Tablet{backpackWithLaptor}
	backpackWithLaptorAndTabletAndWater := &BottleOfWater{backpackWithLaptorAndTablet}

	fmt.Println("Weight of backpack with Laptop, Tablet and water are", backpackWithLaptorAndTabletAndWater.getWeight())
}
