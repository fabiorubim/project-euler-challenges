package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Choose an algorithm to execute:")
		fmt.Println("1. Multiples of 3 and 5")
		fmt.Println("2. Even Fibonacci Numbers")
		fmt.Println("0. Exit")

		var choice int
		fmt.Println("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("error on choice:", err)
			os.Exit(1)
		}
		switch choice {
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		case 1:
			multiples_3_5()
		case 2:
			even_fibonacci_numbers()
		default:
			fmt.Println("Invalid option. Try again!")
		}
		fmt.Println()
	}
}
