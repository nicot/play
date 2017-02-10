package guru

import "testing"

func TestEvalRec1(t *testing.T) {
	e := Parse("(+ 1 2)")
	v := Value{Number, 3}

	c := EvalRec(e)

	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEvalRec2(t *testing.T) {
	e := Parse("(+ 2 (+ 8 (+ 2 3)))")
	v := Value{Number, 15}

	c := EvalRec(e)

	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEvalRec3(t *testing.T) {
	e := Parse("(+ (+ 1 2) 2)")
	v := Value{Number, 5}

	c := EvalRec(e)

	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEvalRec4(t *testing.T) {
	e := Parse("(+ (+ 1 2) (+ 7 8))")
	v := Value{Number, 18}

	c := EvalRec(e)

	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}

func TestEval1(t *testing.T) {
	t.Skip()
	e := Parse("(+ 1 2)")
	v := Value{Number, 18}

	c := Eval(e)

	if c != v {
		t.Errorf("Expected %v to equal %v", c, v)
	}
}
