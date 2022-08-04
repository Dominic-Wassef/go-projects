package main

import "fmt"

type shape interface {
	getArea() float64
}

type squaure struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s squaure) getArea() float64 {
	tot := s.sideLength * s.sideLength
	return tot
}

func (t triangle) getArea() float64 {
	total := (0.5 * t.base * t.height)
	return total
}

func printArea(t triangle, s squaure) {
	fmt.Println("Square:", s.getArea())
	fmt.Println("Triangle:", t.getArea())
}

func main() {
	tin := triangle{
		height: 24.393,
		base:   49.293,
	}
	sin := squaure{
		sideLength: 32.394,
	}
	printArea(tin, sin)
}
