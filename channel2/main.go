package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	names := []string{"Annisa", "Nadia", "Neyla"}
	ch := make(chan string)

	for _, name := range names {
		go func(data string) {
			message := fmt.Sprintf("Hello, %s!", data)
			ch <- message
		}(name)
	}

	for i := 0; i < len(names); i++ {
		printMessage(ch)
	}
}

func printMessage(ch chan string) {
	fmt.Println("Message:", <-ch)
}
