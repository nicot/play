package main

import "math"

type Range struct {
	start int
	end int
}

func invert_ranges(ranges [][]]Range) []Range {
	initial := []Range{
			Range{
			math.MinInt64,
			math.MaxInt64,
		}
	}

	currentRange := initial

	for _, rv := range ranges {
		currentRange = invert_range(currentRange, rv)
	}

	return currentRange
}

func invert_range(range1 []Range, range2 []Range) []Range {
}
