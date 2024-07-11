package main

import "fmt"

type triangle struct {
	height int
	weight int
}

type square struct {
	side int
}

type shape interface {
	getArea() int
}

func (t triangle) getArea() int {
	tarea := 0.5 * float64(t.height) * float64(t.weight)

	return int(tarea)
}

func (s square) getArea() int {
	sarea := s.side * s.side

	return int(sarea)
}

func printArea(b shape) {
	fmt.Println(b.getArea())
}

func main() {
	var t triangle
	var s square

	t.height = 10
	t.weight = 10

	s.side = 10

	printArea(t)
	printArea(s);

}