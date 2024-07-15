package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s student) greet() {
	fmt.Printf("Hello, %s!\n", s.name)
}

func (s student) getName() string {
	return s.name
}

func main() {
	student1 := student{name: "Nadia", grade: 9}

	student1.greet()
	fmt.Println(student1.getName())
}
