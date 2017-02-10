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
	e := Parse("(+ 2 (+ 8 (+ 2 3)))")
	v := Value{TNumber, 15}

	c := Eval(e)
	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEval3(t *testing.T) {
	e := Parse("(+ (+ 1 2) 2)")
	v := Value{TNumber, 5}

	c := Eval(e)
	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEval4(t *testing.T) {
	e := Parse("(+ (+ 1 2) (+ 7 8))")
	v := Value{TNumber, 18}

	c := Eval(e)
	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}
