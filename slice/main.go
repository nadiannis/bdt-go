package main

import "fmt"

func main() {
	var arr [3]int
	var slice []string

	arr[0] = 1
	slice = append(slice, "nadia")

	fmt.Println(arr, len(arr))
	fmt.Println(slice, len(slice))
}
