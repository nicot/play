package main

import "math"

type Range struct {
	start int
	end int
}

func invert_ranges(ranges [][]Range) []Range {
	currentRange := []Range{
		start: math.MinInt64,
		end: math.MaxInt64,
	}

	for _, rv := range ranges {
		currentRange = invert_range(currentRange, rv)
	}

	return currentRange
}

func invert_range(range1 []Range, range2 []Range) []Range {
	if len(range1) == 0 {
		return range2
	}
	if len(range2) == 0 {
		return range1
	}

	next := make([]Range{}, 0)
	currentEnd := math.MaxInt64
	i := 0
	j := 0
	for {
		// todo length checks
		var smaller Range
		if range1[0].start < range2[0].start {
			smaller = range1[i]
			i++
		} else {
			smaller = range2[j]
			j++
		}
		if smaller.end < currentEnd {
			next[len(next) - 1].end = smaller.end
		}

		// actually do the inverting
		next = append(next, smaller)
		currentEnd = smaller.end
	}
}
