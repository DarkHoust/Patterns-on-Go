package main

import (
	"fmt"
	"math"
)

type Hole struct {
	Radius float32
}

type RoundPeg struct {
	Radius float32
}

type SquarePeg struct {
	Width float32
}

type IHole interface {
	GetRadius() float32
	IsFits(rp RoundPeg) bool
}

type IRoundPeg interface {
	GetRadius() float32
}

type ISquarePeg interface {
	GetWidth() float32
}
type RoundHoleAdapter struct {
	SquarePeg ISquarePeg
}

func (a RoundHoleAdapter) GetRadius() float32 {
	// Assuming the effective radius for a square peg is half of the diagonal
	diagonal := math.Sqrt(float64(a.SquarePeg.GetWidth()*a.SquarePeg.GetWidth() + a.SquarePeg.GetWidth()*a.SquarePeg.GetWidth()))
	return float32(diagonal / 2)
}

func HoleIsFits(h Hole, rp RoundPeg) bool {
	return h.Radius >= rp.Radius
}

func (sp SquarePeg) GetWidth() float32 {
	return sp.Width
}

func main() {
	roundHole := Hole{Radius: 5.0}

	roundPeg := RoundPeg{Radius: 3.0}

	fmt.Println("Round Peg fits into Round Hole:", HoleIsFits(roundHole, roundPeg))

	squarePeg := SquarePeg{Width: 4.0}

	adapter := RoundHoleAdapter{SquarePeg: squarePeg}

	fmt.Println("Square Peg fits into Round Hole (via Adapter):", HoleIsFits(roundHole, RoundPeg{Radius: adapter.GetRadius()}))
}
