package main

import (
	"fmt"
	"math"
)

// Largest prime factor
func largest_prime_factor() {
	fmt.Println("largest prime factor")
	// number := 30
	number := 600851475143
	maxNumber := int(math.Sqrt(float64(number)))
	intNumbers := make([]int, number)
	for n := 2; n < number; n++ {
		intNumbers[n] = n
	}
	newIntNumbers := []int{}
	for n := 2; n <= maxNumber; n++ {
		for i := n + 1; i < number; i++ {
			if intNumbers[i]%n == 0 {
				newIntNumbers = append(intNumbers[:i], intNumbers[i+1:]...)
			}
		}
	}

	fmt.Println(newIntNumbers)
}
