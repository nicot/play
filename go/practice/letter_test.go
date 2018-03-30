package practice

import "testing"

func withinOneEdit(a string, b string) bool {
	var s, l string
	if len(a) < len(b) {
		s = a
		l = b
	} else {
		s = b
		l = a
	}

	if len(l)-len(s) > 1 {
		return false
	}

	edited := false

	for i := 0; i < len(l); i++ {
		if i == len(s) {
			// if we have successfully reached the last element of s with no edits, we succeeded!
			return true
		}

		cs := s[i]
		cl := l[i]

		if cs == cl {
			continue
		}

		if edited {
			return false
		}
		edited = true
		if len(l) == len(s) {
			// just skip this index, count this letter as a replace.
			continue
		}

		// Otherwise add a one char string to the beginning of the short string, and continue on the next index.
		// O(number of bytes in s)
		s = " " + s
	}
	return true
}

func TestFoo(t *testing.T) {
	c := withinOneEdit("foo", "boo")
	t.Error(c)

	b := withinOneEdit("foo", "bo")
	t.Error(b)

	t.Error(withinOneEdit("baz", "baf"))
	t.Error(withinOneEdit("foo", "fo"))
	t.Error(withinOneEdit("fo", "foo"))
	t.Error(withinOneEdit("oo", "foo"))
	t.Error(withinOneEdit("fooo", "fo"))
}
