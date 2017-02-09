package guru

import "testing"

func TestParse1(t *testing.T) {
	tokens := Lex("(+ 1 2)")
	exTree := Node{nil, []Node{
		{&Token{Symbol, "+"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
	}}

	tree := Parse(tokens)
	if !tree.Equals(exTree) {
		t.Errorf("Expected tree %v to be %v", tree, exTree)
	}
}

func TestParse2(t *testing.T) {
	tokens := Lex("(+ 1 2 (+ 2 3))")
	exTree := Node{nil, []Node{
		{&Token{Symbol, "+"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
		{nil, []Node{
			{&Token{Symbol, "+"}, nil},
			{&Token{Number, 2}, nil},
			{&Token{Number, 3}, nil},
		}},
	}}

	tree := Parse(tokens)
	if !tree.Equals(exTree) {
		t.Errorf("Expected tree %v to be %v", tree, exTree)
	}
}

func TestTreeEquals(t *testing.T) {
	tree1 := Node{nil, []Node{
		{&Token{Symbol, "+"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
	}}
	tree2 := Node{nil, []Node{
		{&Token{Symbol, "+"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
	}}
	f := tree1.Equals(tree2)
	if !f {
		t.Errorf("%v should equal %v", tree1, tree2)
	}

	tree3 := Node{nil, []Node{
		{&Token{Symbol, "-"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
	}}

	f = tree1.Equals(tree3)
	if f {
		t.Errorf("%v shouldn't equal %v", tree1, tree3)
	}
}
