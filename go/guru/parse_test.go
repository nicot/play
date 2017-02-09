package guru

import "testing"

func TestParse1(t *testing.T) {
	tokens := Read("(+ 1 2)")
	exTree := Node{nil, []Node{
		{&Token{Symbol, "+"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
	}}

	tree := Parse(tokens)
	t.Error(tree.Equals(exTree))
	t.Error(tree)
	t.Error(exTree)
}

func TestParse2(t *testing.T) {
	tokens := Read("(+ 1 2 (+ 2 3))")
	exTree := Node{nil, []Node{
		{&Token{Symbol, "+"}, nil},
		{&Token{Number, 1}, nil},
		{&Token{Number, 2}, nil},
	}}

	tree := Parse(tokens)
	t.Error(tree.Equals(exTree))
	t.Error(tree)
	t.Error(exTree)
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
