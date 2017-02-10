package guru

import "testing"

func TestParse1(t *testing.T) {
	exTree := Node{nil, []Node{
		{&Value{Symbol, "+"}, nil},
		{&Value{Number, 1}, nil},
		{&Value{Number, 2}, nil},
	}}

	tree := Parse("(+ 1 2)")
	if !tree.Equals(exTree) {
		t.Errorf("Expected tree %v to be %v", tree, exTree)
	}
}

func TestParse2(t *testing.T) {
	exTree := Node{nil, []Node{
		{&Value{Symbol, "+"}, nil},
		{&Value{Number, 1}, nil},
		{&Value{Number, 2}, nil},
		{nil, []Node{
			{&Value{Symbol, "+"}, nil},
			{&Value{Number, 2}, nil},
			{&Value{Number, 3}, nil},
		}},
	}}

	tree := Parse("(+ 1 2 (+ 2 3))")
	if !tree.Equals(exTree) {
		t.Errorf("Expected tree %v to be %v", tree, exTree)
	}
}

func TestTreeEquals(t *testing.T) {
	tree1 := Node{nil, []Node{
		{&Value{Symbol, "+"}, nil},
		{&Value{Number, 1}, nil},
		{&Value{Number, 2}, nil},
	}}
	tree2 := Node{nil, []Node{
		{&Value{Symbol, "+"}, nil},
		{&Value{Number, 1}, nil},
		{&Value{Number, 2}, nil},
	}}
	f := tree1.Equals(tree2)
	if !f {
		t.Errorf("%v should equal %v", tree1, tree2)
	}

	tree3 := Node{nil, []Node{
		{&Value{Symbol, "-"}, nil},
		{&Value{Number, 1}, nil},
		{&Value{Number, 2}, nil},
	}}

	f = tree1.Equals(tree3)
	if f {
		t.Errorf("%v shouldn't equal %v", tree1, tree3)
	}
}

func TestParse3(t *testing.T) {
	exTree := Node{nil, []Node{
		{&Value{Symbol, "+"}, nil},
		{nil, []Node{
			{&Value{Symbol, "+"}, nil},
			{&Value{Number, 2}, nil},
			{&Value{Number, 3}, nil},
		}},
		{&Value{Number, 2}, nil},
	}}

	tree := Parse("(+ (+ 2 3) 2)")
	if !tree.Equals(exTree) {
		t.Errorf("Expected tree %v to be %v", tree, exTree)
	}
}
