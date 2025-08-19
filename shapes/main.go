package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{
		base:   5,
		height: 10,
	}
	s := square{
		side: 4,
	}

	printArea(t)
	printArea(s)

}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}
