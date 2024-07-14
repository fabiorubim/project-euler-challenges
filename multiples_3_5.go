package main

import "fmt"

// Multiples 3 or 5
func multiples_3_5() {
	fmt.Println("Multiple of 3 or 5 Algorithm")
	multiples := []int{}
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
		}
	}

	sum := 0
	for _, value := range multiples {
		sum += value
	}

	fmt.Println(sum)
}
