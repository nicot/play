package guru

import (
	"testing"
)

func TestLex1(t *testing.T) {
	s := "(+ 1 2)"
	tokens := Lex(s)
	expectedTokens := []Token{
		{Delimiter, "open"},
		{TSymbol, "+"},
		{TNumber, 1},
		{TNumber, 2},
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
		{TSymbol, "append"},
		{TString, "foo"},
		{TNumber, 2},
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
		{TSymbol, "append"},
		{TString, "foo"},
		{Delimiter, "open"},
		{TSymbol, "+"},
		{TNumber, 12},
		{TNumber, 2},
		{Delimiter, "close"},
		{Delimiter, "close"},
	}
	for i, v := range expectedTokens {
		if tokens[i] != v {
			t.Errorf("Expected token %v in position %d, got %v", v, i, tokens[i])
		}
	}
}
