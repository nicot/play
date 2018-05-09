package main

type Heap struct {
	contents []int
	greaterThan func(a, b int) bool
}

func MinHeap() Heap {
	return Heap{
		greaterThan: func (a, b int) bool {
			return a < b
		},
	}
}

func MaxHeap() Heap {
	return Heap{
		greaterThan: func (a, b int) bool {
			return a > b
		},
	}
}


func (h *Heap) Insert(n int) {
	if h.contents == nil {
		h.contents = make([]int, 1)
	}
	h.contents = append(h.contents, n)

	i := len(h.contents) - 1
	v := h.contents[i]

	for i != 0 {
		pi := (i - 1) / 2
		parent := h.contents[pi]
		if h.greaterThan(parent, v) {
			break
		}
		h.contents[i] = parent
		h.contents[pi] = v
		i = pi
	}
}

func (h *Heap) FindMax() int {
	return h.contents[0]
}

func (h *Heap) DeleteMax() {
	if len(h.contents) == 0 {
		return
	}

	h.contents[0] = h.contents[len(h.contents) - 1]

	h.contents = h.contents[:len(h.contents) - 1]

	i := 0

	for {
		v := h.contents[i]
		left := i*2 + 1
		right := i*2 + 2
		nextSwap := i

		if left < len(h.contents) && h.greaterThan(h.contents[left], h.contents[nextSwap]) {
			nextSwap = left
		}
		if right < len(h.contents) && h.greaterThan(h.contents[right], h.contents[nextSwap]) {
			nextSwap = right
		}

		if nextSwap == i {
			// i is smaller than it's children, so we're done
			break
		}
		h.contents[i] = h.contents[nextSwap]
		h.contents[nextSwap] = v
		i = nextSwap
	}
}

func (h *Heap) Verify() bool {
	for i := 0; i*2 + 2 < len(h.contents); i++ {
		v := h.contents[i]

		c1, c2 := h.contents[2*i+1], h.contents[2*i+2]

		if h.greaterThan(v, c1) || h.greaterThan(v, c2) {
			return false
		}
	}
	return true
}

func (h *Heap) Size() int {
	return len(h.contents)
}

func runningMedian(n int, smaller Heap, bigger Heap) int {
	smallerMid := smaller.FindMax()
	biggerMid := bigger.FindMax()

	smallerSize := smaller.Size()
	biggerSize := bigger.Size()

	if bigger.Size() > smaller.Size() {
		if n > biggerMid {
			bigger.DeleteMax()
			bigger.Insert(n)
			smaller.Insert(biggerMid)
		} else {
			smaller.Insert(n)
		}
	} else {
		if n < smallerMid {
			smaller.DeleteMax()
			smaller.Insert(n)
			bigger.Insert(smallerMid)
		} else {
			bigger.Insert(n)
		}
	}

	if bigger.Size() == smaller.Size() {
		return (bigger.FindMax() + smaller.FindMax()) / 2
	}
	if bigger.Size() > smaller.Size {
		return bigger.FindMax()
	}
	return smaller.FindMax()c
}
