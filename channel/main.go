package main

import "fmt"

func main() {
	ch := make(chan string)
	defer close(ch)

	go getMessage(ch)

	message := <-ch
	fmt.Println(message)
}

func getMessage(ch chan string) {
	ch <- "Hello, world!"
}
