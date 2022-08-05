package main

import "fmt"

type shape interface {
	getArea() float64
}

type rectangle struct {
	base   float64
	height float64
}

type triangle struct {
	sideLength float64
}

func (r rectangle) getArea() float64 {
	return 0.5 * r.base * r.height
}

func (t triangle) getArea() float64 {
	return t.sideLength * t.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	re := rectangle{
		base:   23,
		height: 32,
	}
	tri := triangle{
		sideLength: 21,
	}
	printArea(re)
	printArea(tri)
}
