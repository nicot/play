package practice

import (
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	l := mklist(10)
	quickSortRec(l)
	x := l[0]
	for i, v := range l[1:] {
		if v < x {
			t.Error("Sort failed", v, x, i, l)
		}
	}
}

func TestBenchSort(t *testing.T) {
	t.Skip()
	l := mklist(1000 * 1000 * 10)
	start := time.Now()
	quickSortRec(l)
	t.Error(time.Since(start))
}

func TestBenchAppend(t *testing.T) {
	t.Skip()
	n := 400
	s := make([]int, n)
	f := make([]int, 2, n+3)
	start := time.Now()
	l := append(f, s...)
	t.Error(time.Since(start))
	if cap(f) != cap(l) {
		t.Error(cap(f), cap(l))
	}
}

func TestMerge(t *testing.T) {
	l := mklist(100)
	l = mergeSort(l)
	x := l[0]
	for i, v := range l[1:] {
		if v < x {
			t.Error("Sort failed", v, x, i, l)
		}
	}
}

func TestBenchMergeSort(t *testing.T) {
	t.Skip()
	l := mklist(1000 * 1000 * 10)
	start := time.Now()
	mergeSort(l)
	t.Error(time.Since(start))
}
