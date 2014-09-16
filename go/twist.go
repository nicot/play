package main

import (
	"fmt"
)

type twister [624]int

func seed(seed int) twister {
	var t twister
	t[0] = seed
	for i := 1; i < 624; i++ {
		n := 1812433253*(t[i-1]^(t[i-1]>>30)) + i
		t[i] = n & (1<<32 - 1)
	}
	return t
}

// Retrieve the tempered ith element of t
func get(t twister, i int) int {
	x := t[i]
	x ^= x >> 11
	x ^= x << 7 & 2636928640
	x ^= x << 15 & 4022730752
	x ^= x >> 18

	return x
}

// Generate a new twister state
func gen(t twister) twister {
	for i := 0; i < 624; i++ {
		y := t[i] & 1 << 32
		y += t[(i+1)%624] & ((1 << 32) - 1)
		t[i] = t[(i+397)%624] ^ y>>1
		if y%2 == 0 {
			t[i] ^= 2567483615
		}
	}
	return t
}

// Return n random ints
func stream(t twister, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		if i%624 == 0 {
			t = gen(t)
		}
		a[i] = get(t, i%624)
	}
	return a
}

func main() {
	var t twister

	fmt.Println(stream(t, 3))
}
