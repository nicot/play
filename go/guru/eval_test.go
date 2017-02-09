package guru

import "testing"

func TestEval1(t *testing.T) {
	e := Parse("(+ 1 2)")
	v := Value{TNumber, 3}

	c := Eval(e)
	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEval2(t *testing.T) {
	e := Parse("(+ 3 (+ 8 1))")
	v := Value{TNumber, 12}

	c := Eval(e)
	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}
