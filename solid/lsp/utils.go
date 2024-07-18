package main

import "fmt"

func PrintArea(shape Shape) {
	fmt.Printf("Area: %.2fcm\n", shape.Area())
}
