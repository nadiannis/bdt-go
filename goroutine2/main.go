package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Num of CPU core available:", runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	now := time.Now()

	wg.Add(2)
	go print(100, "Hello, World!", now, &wg)
	print(100, "I'm learning Go", now, &wg)

	wg.Wait()
	fmt.Println("End of main()")
}

func print(n int, message string, startTime time.Time, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		dur := time.Since(startTime)
		fmt.Printf("%d | process time: %dns | %s\n", (i + 1), dur.Nanoseconds(), message)
	}
	wg.Done()
}
