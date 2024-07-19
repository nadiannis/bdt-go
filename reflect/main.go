package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 22
	reflectValue := reflect.ValueOf(number)

	fmt.Printf("number: %+v\n", number)
	fmt.Printf("reflectValue: %+v\n", reflectValue)

	fmt.Printf("type: %+v\n", reflectValue.Type())
	fmt.Printf("kind: %+v\n", reflectValue.Kind())
	fmt.Printf("interface: %+v\n", reflectValue.Interface())
	fmt.Println()

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("The variable is a type of", reflectValue.Type())
		fmt.Println("Here's the value:", reflectValue.Int())
	} else {
		fmt.Println("The variable is a type of", reflect.TypeOf(number))
	}

	if reflectValue.Kind() == reflect.Int {
		fmt.Printf("interface.(int): %+v\n", reflectValue.Interface().(int))
	}
}
