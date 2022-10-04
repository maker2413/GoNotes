package main

func Sum(x []int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}

	return sum
}

func SumAll(x ...[]int) []int {
	var sums []int

	for _, y := range x {
		sums = append(sums, Sum(y))
	}
	
	return sums
}

func SumAllTails(x ...[]int) []int {
	var sums []int

	for _, y := range x {
		if len(y) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(y[1:]))
		}
	}
	
	return sums
}
