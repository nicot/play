package main

import (
	"fmt"
	"sort"
	"strconv"
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
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			arr[j] = readInt()
		}
		result := isTroubled(arr)
		fmt.Printf("Case #%d: %s\n", i + 1, result)
	}
}

type Pair struct {
	value int
	index int
}

type PairSorter struct {
	pairs []Pair
}

func (p PairSorter) Len() int {
	return len(p.pairs)
}

func (p PairSorter) Less(i, j int) bool {
	return p.pairs[i].value < p.pairs[j].value
}

func (p PairSorter) Swap(i, j int) {
	tmp := p.pairs[i]
	p.pairs[i] = p.pairs[j]
	p.pairs[j] = tmp
}

func isTroubled(arr []int) string {
	pairs := make([]Pair, len(arr))
	for i, n := range arr {
		// turn them into pairs of {index, value}
		pairs[i] = Pair{value: n, index: i}
	}
	p := PairSorter{pairs}
			// sort based on value
	sort.Sort(p)

	n := solve(p.pairs)
	if n < 0 {
		return "OK"
	} else {
		return strconv.Itoa(n)
	}
}

// Return first bad index or -1
func solve(pairList []Pair) int {
	for index, pair := range pairList {
		if index % 2 == pair.index % 2 {
			continue
		}

		for i := index + 1; i < len(pairList); i++ {
			candidate := pairList[i]
			if candidate.value != pair.value {
				return index
			}
			if candidate.index % 2 != pair.index % 2 {
				pairList[index] = candidate
				pairList[i] = pair
				break
			}
		}
	}
	return -1
}
