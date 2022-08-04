package main

import "fmt"

type shape interface {
	getShape() float64
}
type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t triangle) getShape() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getShape() float64 {
	return s.sideLength * s.sideLength
}

func printShape(s shape) {
	fmt.Println(s.getShape())
}

func main() {
	tri := triangle{
		base:   43,
		height: 32,
	}
	squ := square{
		sideLength: 11,
	}
	printShape(tri)
	printShape(squ)
}
