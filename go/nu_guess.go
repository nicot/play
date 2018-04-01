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
	// read the input from stdin
	// A, B A < P <= B
	// N guess
	// goal is to gues P,
	var t int
	_, err := fmt.Scanf("%d", &t)
	if err != nil {
		panic(err)
	}

	for i := 0; i < t; i++ {
		solveCase()
	}
}

func solveCase() {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		panic(err)
	}
	a++

	var n int
	_, err = fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	for {
		median := (a + b) / 2
		_, err :=  fmt.Println(median)
		if err != nil {
			panic(err)
		}

		var input string
		_, err = fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}

		if input == "CORRECT" {
			return
		}
		if input == "WRONG_ANSWER" {
			panic(input)
		}
		if input == "TOO_SMALL" {
			a = median + 1
		} else if input == "TOO_BIG" {
			b = median - 1
		} else {
			panic(input)
		}
	}
}