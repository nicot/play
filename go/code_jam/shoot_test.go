package main

import "testing"

func TestFoo(t *testing.T) {
	t.Error(calcDamage([]byte("CSCCSCSSS")))
	t.Error(solve(6, []byte("SCCSSC")))
}