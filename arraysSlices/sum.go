package main

import "fmt"

func Sum(numbers [5]int) int {
	var sum int
	// range lets you iterate over an array. On each iteration,
	// range returns two values - the index and the value.
	// We are choosing to ignore the index value by using _ blank identifier.
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(Sum(numbers))
}
