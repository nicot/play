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
	n := readInt()
	if n <= 0 {
		panic("must be positive")
	}
	elements := make([]int, n)
	weights := make([]int, n)
	for i := 0; i < n; i++ {
		elements[i] = readInt()
	}
	for i := 0; i < n; i++ {
		weights[i] = readInt()
	}

	fmt.Printf("%.1f", calcMean(elements, weights))
}

func calcMean(elements []int, weights []int) float64 {
	var divisor, numerator int
	for i := range elements {
		divisor += weights[i]
		numerator += weights[i] * elements[i]
	}

	return float64(numerator) / float64(divisor)
}