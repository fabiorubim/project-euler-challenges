package main

import "fmt"

// Even Fibonacci Numbers
func even_fibonacci_numbers() {
	fmt.Println("Executing Even Fibonacci Numbers Algorithm")
	a := 0
	sum := 0
	for {
		a++
		fibVal := fibonacci(a)
		if fibVal <= 4000000 {
			if fibVal%2 == 0 {
				sum += fibVal
			}
		} else {
			break
		}
	}
	fmt.Println(sum)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
