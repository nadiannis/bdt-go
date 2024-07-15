package main

import "fmt"

func main() {
	var users map[string]int = map[string]int{}

	users["nadia"] = 24
	users["john"] = 45
	users["taylor"] = 40

	fmt.Println(users)

	for name, age := range users {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
