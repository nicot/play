package guru

import (
	"testing"
)

func TestLex1(t *testing.T) {
	s := "(+ 1 2)"
	tokens := Lex(s)
	expectedTokens := []Token{
		{Delimiter, "open"},
		{Symbol, "+"},
		{Number, 1},
		{Number, 2},
		{Delimiter, "close"},
	}
	for i, v := range expectedTokens {
		if tokens[i] != v {
			t.Errorf("Expected token %v in position %d, got %v", v, i, tokens[i])
		}
	}
}

func TestLex2(t *testing.T) {
	s := `(append "foo" 2)`
	tokens := Lex(s)
	expectedTokens := []Token{
		{Delimiter, "open"},
		{Symbol, "append"},
		{String, "foo"},
		{Number, 2},
		{Delimiter, "close"},
	}
	for i, v := range expectedTokens {
		if tokens[i] != v {
			t.Errorf("Expected token %v in position %d, got %v", v, i, tokens[i])
		}
	}
}

func TestLex3(t *testing.T) {
	s := `(append "foo" (+ 12 2))`
	tokens := Lex(s)
	expectedTokens := []Token{
		{Delimiter, "open"},
		{Symbol, "append"},
		{String, "foo"},
		{Delimiter, "open"},
		{Symbol, "+"},
		{Number, 12},
		{Number, 2},
		{Delimiter, "close"},
		{Delimiter, "close"},
	}
	for i, v := range expectedTokens {
		if tokens[i] != v {
			t.Errorf("Expected token %v in position %d, got %v", v, i, tokens[i])
		}
	}
}
