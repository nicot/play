package skip

import (
	"bytes"
	"math/rand"
	"testing"
)

func setup() (Skiplist, map[int][]byte) {
	size := 10
	m := make(map[int][]byte, size)
	s := Skiplist{}
	for i := 0; i < size; i++ {
		k := rand.Int()
		v := make([]byte, 10)
		rand.Read(v)
		s.Set(k, v)
		m[k] = v
	}
	return s, m
}

func TestSet(t *testing.T) {
}

func TestGet(t *testing.T) {
	s, m := setup()
	for k, v := range m {
		got := s.Get(k).([]byte)
		if !bytes.Equal(got, v) {
			t.Errorf("expected %v, got %v", got, v)
		}
	}
}
