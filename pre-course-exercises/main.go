package main

import (
	"fmt"
	"math"
)

type shape struct {
	getArea float64
}

type square struct {
	width int
}

type circle struct {
	radius int
}

func (c circle) getArea() float64 {
	return (math.Pi * float64(c.radius*c.radius))
}

func (s square) getArea() float64 {
	return float64(s.width * s.width)
}

func main() {
	mySquare := square{2}
	myCircle := circle{2}

	fmt.Println(mySquare.getArea())
	fmt.Println(myCircle.getArea())
}
