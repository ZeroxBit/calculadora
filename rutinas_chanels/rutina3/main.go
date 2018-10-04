package main

import (
	"fmt"
	"time"
)

func main() {
	go animation(100 * time.Millisecond)
	const n = 43
	result := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d \n", n, result)
}

func animation(retraso time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}

}

func fib(x int) int { // fibonacci
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
