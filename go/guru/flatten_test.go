package guru

import "testing"

func TestFlatten1(t *testing.T) {

}

func TestFlatten2(t *testing.T) {
	e := Parse("(if true 1 2)")

	f := Flatten(e)

	v := []FlatExpr{}

	if len(f) != len(v) {
		t.Errorf("Got %v of length %d, expected %v of %d", f, v, len(f), len(v))
	}
}

func TestFlattenLazy(t *testing.T) {
	e := Parse("(if true (print 2) (print 3))")

	c := Flatten(e)

	v := ByteCode{
		FlatExpr{
			Value{},
			[]Value{},
		},
	}

	if !c.Equals(v) {
		t.Errorf("Expected %v, got %v", c, v)
	}
}
