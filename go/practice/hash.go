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

// Can only hold 2^56? bits of items
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
	i := 0
	for {
		h1 := (h.hash(v1.key) + i) % h.cap
		// Robin hood!
		v2 := h.vals[h1]
		if v2 == nil {
			h.len++
			h.vals[h1] = v1
			break
		}
		if v2.key == v1.key {
			h.vals[h1] = v1
			break
		}

		h2 := h.hash(v2.key)

		d := (h1 - i) % h.cap
		if d < 0 {
			d += h.cap
		}
		if d < h2 {
			h.vals[h1] = v1
			v1 = v2
		}
		i++
	}
	fmt.Println(key, value, h)
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

}
