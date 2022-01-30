package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

func main() {
	t1 := triangle{height: 10, base: 10}
	s1 := square{side: 10}

	printArea(t1)
	printArea(s1)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
