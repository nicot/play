package practice

import (
	"math/rand"
)

func mklist(n int) []int64 {
	var l []int64
	for i := 0; i < n; i++ {
		v := rand.Int63()
		l = append(l, v)
	}
	return l
}

func mergeSort(l []int64) []int64 {
	if len(l) == 1 {
		return l
	}
	split := len(l) / 2
	left := mergeSort(l[:split])
	right := mergeSort(l[split:])
	leftIndex := 0
	rightIndex := 0
	end := make([]int64, 0, len(l))
	for i := 0; i < len(l); i++ {
		if len(left) == leftIndex {
			end = append(end, right...)
			break
		}
		if len(right) == rightIndex {
			end = append(end, left...)
			break
		}
		if left[leftIndex] < right[rightIndex] {
			end = append(end, left[leftIndex])
			leftIndex++
		} else {
			end = append(end, right[rightIndex])
			rightIndex++
		}
	}
	return end
}

func quickSortRec(l []int64) {
	if len(l) < 2 {
		return
	}

	pivot := 0 // Can also choose randomly
	leftIndex := 1
	rightIndex := len(l)
	for leftIndex < rightIndex {
		if l[leftIndex] < l[pivot] {
			leftIndex++
		} else {
			rightIndex--
			l[rightIndex], l[leftIndex] = l[leftIndex], l[rightIndex]
		}
	}
	leftIndex--
	l[pivot], l[leftIndex] = l[leftIndex], l[pivot]
	quickSortRec(l[:leftIndex])
	quickSortRec(l[rightIndex:])
}
