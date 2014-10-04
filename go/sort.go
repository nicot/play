package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func verifySorted(l []int) bool {
	for i := 1; i < len(l); i++ {
		if l[i] > l[i+1] {
			return false
		}
	}
	return true
}

func verifyRev(l []int) (bool, int, int) {
	start := -1
	end := -1
	for i := range l {
		if l[i] > l[i+1] {
			if start == -1 {
				start = i
			}
		} else if start > -1 {
			end = i
			s = sorted(l[i:])
			return s, start, end
		}

	}
}

func verifySwap(l []int) (bool, int, int) {
	start := -1
	for i := range l {
		if l[i] > l[i+1] && start == -1 {
			start = i
		}
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	check(err)
	l := make([]int, n)
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	check(err)
	for i, n := range strings.Split(s, " ") {
		l[i], err = strconv.Atoi(n)
		check(err)
	}
	check(err)
	fmt.Println(l)
	var start, end int

	sorted := verifySorted(l)
	if sorted {
		fmt.Println("yes")
		return
	}

	swap, start, end := verifySwap(l)
	if swap {
		fmt.Println("yes")
		fmt.Println("swap", start, end)
		return
	}

	rev, start, end := verifyRev(l)
	if rev {
		fmt.Println("yes")
		fmt.Println("reverse", start, end)
		return
	}

	fmt.Println("no")
}
