package main

import (
	"fmt"
	"math"
)

func sum_square_difference() {
	total_natural_numbers_at_end := sum_of_natural_number()
	total_natural_numbers__at_end_squared := math.Pow(float64(total_natural_numbers_at_end), 2)
	total_each_number_squared := sum_of_natural_number_squared()

	result := int(total_natural_numbers__at_end_squared) - total_each_number_squared
	fmt.Println(result)

}

func sum_of_natural_number() int {
	number := 100
	sum := number * (number + 1) / 2
	return sum
}

func sum_of_natural_number_squared() int {
	number := 100
	sum_squared := number * (number + 1) * (2*number + 1) / 6
	return sum_squared
}
