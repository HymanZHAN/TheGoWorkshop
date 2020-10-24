package main

import "fmt"

// Shape is an interface that defines Area() and Name() method
type Shape interface {
	Area() float64
	Name() string
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) Area() float64 {
	return t.base * t.height / 2
}

func (t triangle) Name() string {
	return "Triangle"
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) Area() float64 {
	return r.width * r.length
}

func (r rectangle) Name() string {
	return "Rectangle"
}

type square struct {
	side float64
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "Square"
}

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	s := square{side: 10}
	printShapeDetails(t, r, s)
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}
