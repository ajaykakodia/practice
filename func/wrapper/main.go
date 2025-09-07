package main

import (
	"fmt"
	"time"
)

func main() {
	timerFunc(doWork)
}

func doWork() {
	for i := 0; i <= 2_000; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	fmt.Println("________________")
}

func timerFunc(f func()) {
	start := time.Now()
	f()
	elapsedTime := time.Since(start)
	fmt.Println(elapsedTime)
}
