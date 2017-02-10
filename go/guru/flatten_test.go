package guru

import "testing"

func TestFlatten1(t *testing.T) {
	e := Parse("(+ 1 2)")

	c := Flatten(e)

	v := FlatExprs{
		[]Value{{Symbol, "+"}, {Number, 1}, {Number, 2}},
	}

	if !c.Equals(v) {
		t.Errorf("Expected %v, got %v", c, v)
	}
}

func TestFlatten2(t *testing.T) {
	t.Skip()
	_ = Parse("(if true 1 2)")

}

func TestFlattenLazy(t *testing.T) {
	t.Skip()
	_ = Parse("(if true (print 2) (print 3))")

}
