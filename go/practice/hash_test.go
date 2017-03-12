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
	h := HashTable{}
	d := make(map[int]int)
	for i := 0; i < 38; i++ {
		k := rand.Intn(40)
		v := rand.Intn(10)
		// k := rand.Int()
		// v := rand.Int()
		h.Set(k, v)
		x := h.Get(k)
		if x.(int) != v {
			t.Error(h, x, v, k)
		}
		d[k] = v
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

}
