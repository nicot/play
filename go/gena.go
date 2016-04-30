package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readIn() []int {
	buf, _ := ioutil.ReadAll(os.Stdin)
	s := string(buf)
	l := strings.Split(s, "\n")
	s = l[1]
	l = strings.Split(s, " ")
	ns := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		ns[i], _ = strconv.Atoi(l[i])
	}
	return ns
}

func main() {
	ns := readIn()
	r := solve(ns)
	fmt.Println(r)
}

func solve(discs []int) int {
	moves := 0
	for !solved(discs) {
		discs = move(discs)
		moves++
	}
	return moves
}

func solved(discs []int) bool {
	for _, v := range discs {
		if v != 1 {
			return false
		}
	}
	return true
}

// hardest: [4 4 4 4 4 4 4 4 4 4 4]
// [4 4] -> [1 4]
func move(discs []int) []int {
	out := biggestOutside(discs)
	// current goal: move out to 1.
	// first: move all discs off out, and off 1.
	// Move out to 1.

	return []int{}
}

func biggestOutsider(discs []int) int {
	for i := len(discs) - 1; i >= 0; i-- {
		if discs[i] != 1 {
			return i
		}
	}
}
