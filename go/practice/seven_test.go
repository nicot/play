package practice

type Line struct {
	intercept int
	slope     int
}

func Intersect(a, b Line) bool {
	if a.intercept == b.intercept {
		return true
	}
	return a.slope == b.slope
}

type Point struct {
	x int
	y int
}

type Square struct {
	center     Point
	sideLength int
}

func HalfSquares(a, b Square) Line {
	// return a line through center of a and center of b
	slope := float64(a.center.x-b.center.x) / float64(a.center.y-b.center.y)
	// y = mx + b
	// 0 = mx + b
	// b = -mx
	intercept := -(float64(a.center.x) * slope)

	return Line{int(intercept), int(slope)}
}
