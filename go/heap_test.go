package main

import "testing"

func TestHeap(t *testing.T) {
	h := MaxHeap()

	h.Insert(1)
	h.Insert(3)
	h.Insert(5)
	h.Insert(8)
	h.Insert(4)

	t.Error(h.Verify())
	t.Error(h)

	h.DeleteMax()
	h.DeleteMax()
	h.DeleteMax()

	t.Error(h)
}