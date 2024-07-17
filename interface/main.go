package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

type Square struct {
	length float64
}

func (s Square) area() float64 {
	return s.length * s.length
}

func (s Square) perimeter() float64 {
	return 4 * s.length
}

func main() {
	square1 := Square{
		length: 3,
	}

	var shape Shape = square1

	fmt.Printf("%+v\n", square1)
	fmt.Println(shape.area())
	fmt.Println(shape.perimeter())
}
