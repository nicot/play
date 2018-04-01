package main

import "fmt"

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", n)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	t := readInt()
	for i := 0; i < t; i++ {
		domax()
	}
}

func domax() {
	n := readInt()
	if n <= 0 {
		panic("n must be positive")
	}
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = readInt()
	}

	maxSubset(array)
	maxSubarray(array)
}

func maxSubset(array []int) {
	var sum int

	for _, v := range array {
		if v > 0 {
			sum += v
		}
	}
	// IDK what to do if all of the numbers are negative.

	fmt.Println(sum)
}

func maxSubarray(array []int) {
	// create an array of prefix sums
	// go through the array, and starting from any positive point, find the biggest

	var max, current int
	for _, v := range array {
		current += v

		if current < 0 {
			current = 0
		}

		if current > max {
			max = current
		}
	}
	// IDK what to do if all of the numbers are negative.
	fmt.Println(sum)
}