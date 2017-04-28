package practice

import (
	"fmt"
)

type HashTable struct {
	vals []*keyVal
	cap  int
	len  int
}

type keyVal struct {
	key int
	val interface{}
}

const maxLoad = .9

func (h *HashTable) hash(key int) int {
	return key % h.cap
}

func (h HashTable) String() string {
	s := fmt.Sprintf("len: %d, cap: %d [", h.len, h.cap)
	for i, v := range h.vals {
		s += fmt.Sprintf("%v", v)
		if i != len(h.vals)-1 {
			s += " "
		}
	}
	return s + "]"
}

func mod(x, y int) int {
	t := x % y
	if t < 0 {
		t += y
	}
	return t
}

// Can only hold 2^53? bits of items
func (h *HashTable) Set(key int, value interface{}) {
	// This can cause it to take more than double load factor if the value hashes to something that already exists
	load := float64(h.len+1) / float64(h.cap)
	if load > maxLoad || h.cap == 0 {
		// grow
		n := h.cap * 2
		if h.cap == 0 {
			n = 1
		}
		vals := make([]*keyVal, n)
		for i := range vals {
			vals[i] = nil
		}
		nh := HashTable{vals: vals, cap: n}
		for _, v := range h.vals {
			if v != nil {
				nh.Set(v.key, v.val)
			}
		}
		*h = nh
	}

	v1 := &keyVal{key, value}
	place := h.hash(v1.key)
	for {
		// Robin hood!
		v2 := h.vals[place]
		if v2 == nil {
			h.len++
			h.vals[place] = v1
			break
		}
		if v2.key == v1.key {
			h.vals[place] = v1
			break
		}

		h1 := mod(h.hash(v1.key)-place, h.cap)
		h2 := mod(h.hash(v2.key)-place, h.cap)
		if h1 > h2 {
			h.vals[place] = v1
			v1 = v2
		}
		place = (place + 1) % h.cap
	}
}

func (h *HashTable) Get(key int) interface{} {
	i := h.hash(key)
	for {
		if h.vals[i] == nil {
			return nil
		}
		v := h.vals[i]
		if v.key == key {
			return v.val
		}
		i = (i + 1) % h.cap
	}
}

func (h *HashTable) Delete(key int) {
	i := h.hash(key)
	for {
		if h.vals[i] == nil {
			break
		}
		v := h.vals[i]
		if v.key == key {
			h.len -= 1
			h.vals[i] = nil
			break
		}
		i = (i + 1) % h.cap
	}
}
