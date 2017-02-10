package guru

import "testing"

func TestFlatten1(t *testing.T) {
	e := Parse("(+ 1 2)")

	c, _ := Flatten(e)

	v := FlatExprs{
		[]Value{{Symbol, "+"}, {Number, 1}, {Number, 2}},
	}
	if !c.Equals(v) {
		t.Errorf("Expected %v, got %v", v, c)
	}
}

func TestFlatten2(t *testing.T) {
	c, _ := Flatten(Parse("(+ 1 (+ 3 4))"))
	v := FlatExprs{
		[]Value{
			{Symbol, "assign"},
			{SSA, 0},
			{Symbol, "+"}, {Number, 3}, {Number, 4},
		},
		[]Value{{Symbol, "+"}, {Number, 1}, {SSA, 0}},
	}

	if !c.Equals(v) {
		t.Errorf("Expected %v, got %v", v, c)
	}
}

func TestFlatten3(t *testing.T) {
	c, _ := Flatten(Parse("(+ (+ 0 1) (+ 5 6) (+ 3 4))"))
	v := FlatExprs{
		[]Value{
			{Symbol, "assign"},
			{SSA, 0},
			{Symbol, "+"}, {Number, 3}, {Number, 4},
		},
		[]Value{{Symbol, "+"}, {Number, 1}, {SSA, 0}},
	}

	if !c.Equals(v) {
		t.Errorf("Expected %v, got %v", v, c)
	}
}

func TestFlattenLazy(t *testing.T) {
	t.Skip()
	_ = Parse("(if true (print 2) (print 3))")

}
