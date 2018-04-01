package main

import (
	"fmt"
	"sort"
)

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
	if n == 0 {
		panic("n should be more than zero")
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	mean(arr)
	median(arr)
	mode(arr)
}

func mean(arr []int) {
	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}
	fmt.Printf("%.1f\n", sum/float64(len(arr)))
}

func median(arr []int) {
	sort.Ints(arr)
	mid := len(arr) / 2
	var v float64
	if len(arr) % 2 == 0 {
		v = float64(arr[mid] + arr[mid-1]) / 2
	} else {
		v = float64(arr[mid])
	}
	fmt.Printf("%.1f\n", v)
}

func mode(arr []int) {
	mode := make(map[int]int)
	var count int
	var mostCommon []int

	for _, v := range arr {
		mode[v]++
		if mode[v] > count {
			mostCommon = []int{v}
			count = mode[v]
		}

		if mode[v] == count {
			mostCommon = append(mostCommon, v)
		}
	}

	sort.Ints(mostCommon)
	fmt.Println(mostCommon[0])
}