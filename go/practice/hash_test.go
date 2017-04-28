package practice

import (
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
	t.Skip()
	h := HashTable{}
	h.Set(0, "foo")
	h.Set(12, "bar")
	h.Set(0, "boof")
	h.Set(3, "baxr")
	h.Set(7, "z")
	h.Set(9, "r")
	i := h.Get(12)
	t.Error(h)
	t.Error(i)
}

func TestSetRand(t *testing.T) {
	t.Skip()
	h := HashTable{}
	d := make(map[int]int)
	for i := 0; i < 400000; i++ {
		k := rand.Int()
		v := rand.Int()
		h.Set(k, v)
		d[k] = v
	}
	for k, v := range d {
		j := h.Get(k)
		if j == nil {
			t.Error(h, k, v)
			t.Error(h.hash(k), h.vals[h.hash(k)])
		}
		if j.(int) != v {
			t.Error(h, k, v, j)
		}
	}
}

func TestDelete(t *testing.T) {
	t.Skip()
	h := HashTable{}
	d := make(map[int]int)
	for i := 0; i < 400; i++ {
		k := rand.Int()
		v := rand.Int()
		h.Set(k, v)
		d[k] = v
	}
	i := 0
	for k := range d {
		if i > 20 {
			break
		}
		i++
		delete(d, k)
		h.Delete(k)
		break
	}
	for k, v := range d {
		j := h.Get(k)
		if j == nil {
			t.Error(h, k, v)
			t.Error(h.hash(k), h.vals[h.hash(k)])
		}
		if j.(int) != v {
			t.Error(h, k, v, j)
		}
	}
}
