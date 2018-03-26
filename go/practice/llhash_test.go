package practice

import (
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	table := Table{}
	m := make(map[int]string)
	for i := 0; i < 200; i++ {
		v := strconv.Itoa(i)
		table.Put(i*33, v)
		m[i*33] = v
	}

	for k, v := range m {
		found := table.Get(k)

		if v != found {
			t.Errorf("Expected %v, got %v", v, found)
		}
	}

	for k := range m {
		delete(m, k)
		table.Delete(k)
		for k, v := range m {
			found := table.Get(k)

			if v != found {
				t.Errorf("Expected %v, got %v", v, found)
			}
		}
	}

	for i := 0; i < 200; i++ {
		v := strconv.Itoa(i)
		table.Put(i*13, v)
		m[i*13] = v
	}

	for k, v := range m {
		found := table.Get(k)

		if v != found {
			t.Errorf("Expected %v, got %v", v, found)
		}
	}
}

// func TestList(t *testing.T) {
// 	var l []int
// 	l = append(l, 1)
// 	t.Error(l[0])
// }
