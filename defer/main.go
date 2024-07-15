package main

import "fmt"

func main() {
	greet()
}

func greet() {
	defer fmt.Println("=== THE END ===")
	fmt.Println("Hello")
	fmt.Println("World")
}