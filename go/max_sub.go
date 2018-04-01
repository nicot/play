package main

import "fmt"

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
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

	fmt.Println(maxSubarray(array), maxSubset(array))
}

func maxSubset(array []int) int {
	var sum int

	max := array[0]

	for _, v := range array {
		if v > 0 {
			sum += v
		}
		if v > max {
			max = v
		}
	}

	if sum > 0 {
		return sum
	} else {
		return max
	}
}

func maxSubarray(array []int) int {
	// create an array of prefix sums
	// go through the array, and starting from any positive point, find the biggest

	var runningMax, current int
	max := array[0]

	for _, v := range array {
		current += v

		if v > max {
			max = v
		}

		if current < 0 {
			current = 0
		}

		if current > runningMax {
			runningMax = current
		}
	}

	if runningMax > 0 {
		return runningMax
	} else {
		return max
	}
}