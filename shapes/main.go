package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	box := square{sideLength: 4}
	tri := triangle{base: 10, height: 10}

	printArea(box)
	printArea(tri)
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(sh shape) {
	fmt.Println("The area is:", sh.getArea())
}
