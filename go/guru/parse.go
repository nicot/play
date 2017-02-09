package guru

import "fmt"

type Node struct {
	// Either
	Token    *Token
	Children []Node
}

func Parse(tokens []Token) Node {
	node, _ := ParsePartial(tokens)
	if node == nil {
		return Node{}
	}
	return *node
}

func ParsePartial(tokens []Token) (*Node, []Token) {
	fmt.Println(tokens)
	token := tokens[0]
	tokens = tokens[1:]
	if token == (Token{Delimiter, "close"}) {
		return nil, tokens
	} else if token == (Token{Delimiter, "open"}) {
		var children []Node
		for {
			node, rest := ParsePartial(tokens)
			tokens = rest
			if node != nil {
				children = append(children, *node)
			}
			if len(tokens) == 0 {
				return &Node{nil, children}, nil
			}
		}
	}
	return &Node{&token, nil}, tokens
}

func (n Node) Equals(n2 Node) bool {
	if n.Token != nil && n2.Token != nil {
		if *n.Token != *n2.Token {
			return false
		}
	}

	if len(n.Children) != len(n2.Children) {
		return false
	}

	if len(n.Children) == 0 {
		return true
	}

	e := true
	for i := range n.Children {
		e = e && n.Children[i].Equals(n2.Children[i])
	}
	return e
}

func (n Node) String() string {
	if n.Token != nil {
		return n.Token.String()
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
