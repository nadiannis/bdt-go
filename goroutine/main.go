package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Num of CPU core available:", runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	now := time.Now()

	go print(100, "Hello, World!", now)
	print(100, "I'm learning Go", now)

	var input string
	fmt.Scanln(&input)

	// time.Sleep(time.Second * 1)
}

func print(n int, message string, startTime time.Time) {
	for i := 0; i < n; i++ {
		dur := time.Since(startTime)
		fmt.Printf("%d | process time: %dns | %s\n", (i + 1), dur.Nanoseconds(), message)
	}
}
