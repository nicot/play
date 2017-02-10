package guru

type Node struct {
	// Either
	Value    *Value
	Children []Node
}

func Parse(input string) Node {
	tokens := Lex(input)
	node, _ := ParsePartial(tokens)
	if node == nil {
		return Node{}
	}
	return *node
}

func ParsePartial(tokens []Token) (*Node, []Token) {
	token := tokens[0]
	tokens = tokens[1:]
	if token == (Token{Delimiter, "close"}) {
		return nil, tokens
	} else if token == (Token{Delimiter, "open"}) {
		var children []Node
		for {
			node, rest := ParsePartial(tokens)
			if node == nil {
				return &Node{nil, children}, rest
			}
			tokens = rest
			children = append(children, *node)
			if len(tokens) == 0 {
				return &Node{nil, children}, nil
			}
		}
	}
	v := TokenToValue(token)
	return &Node{&v, nil}, tokens
}

func (n Node) Equals(n2 Node) bool {
	if n.Value != nil && n2.Value != nil {
		if *n.Value != *n2.Value {
			return false
		}
	}

	if len(n.Children) != len(n2.Children) {
		return false
	}

	if len(n.Children) == 0 {
		return true
	}

	for i := range n.Children {
		if !n.Children[i].Equals(n2.Children[i]) {
			return false
		}
	}
	return true
}

func (n Node) String() string {
	if n.Value != nil {
		return n.Value.String()
	}
	s := "("
	for i, v := range n.Children {
		c := " "
		if i == len(n.Children)-1 {
			c = ""
		}
		s += v.String() + c
	}
	return s + ")"
}
