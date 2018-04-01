package skip

import "math/rand"

type skipList struct {
	Key   int
	Value interface{}
	Next  *skipList
	Down  *skipList
}

// Skiplist is a skiplist
type Skiplist struct {
	Underlying *skipList
}

// Set puts a key into the skiplist
func (s *Skiplist) Set(key int, value interface{}) {
	element := &skipList{
		key,
		value,
		nil,
		nil,
	}
	if s.Underlying == nil {
		s.Underlying = element
		return
	}

	nodes, exists := getNodes(s.Underlying, key, []skipList{})
	if exists {
		nodes[len(nodes)-1].Value = value
		return
	}

	nodes[len(nodes)-1].Next = element

	for i := len(nodes) - 2; i >= 0; i-- {
		if rand.Intn(splitProbability) != 0 {
			return
		}
		prevNode := nodes[i]
		newNode := &skipList{
			key,
			value,
			prevNode.Next,
			&nodes[i+1],
		}
		prevNode.Next = newNode
	}

}

func getNodes(s *skipList, key int, nodes []skipList) ([]skipList, bool) {
	if s == nil {
		return nodes, false
	}

	nodes = append(nodes, *s)

	if s.Key == key {
		return nodes, true
	}
	if s.Next == nil {
		return getNodes(s.Down, key, nodes)
	}
	if key > s.Next.Key {
		return getNodes(s.Next, key, nodes)
	}
	return getNodes(s.Down, key, nodes)
}

const splitProbability = 4

// Get grabs a key from the skiplist
func (s *Skiplist) Get(key int) interface{} {
	v, present := get(s.Underlying, key)

	if !present {
		panic("value not found")
	}
	return v
}

func get(s *skipList, key int) (interface{}, bool) {
	if s == nil {
		return nil, false
	}
	if s.Key == key {
		return s.Value, true
	}
	if s.Next == nil {
		return get(s.Down, key)
	}
	if key > s.Next.Key {
		return get(s.Next, key)
	}
	return get(s.Down, key)
}
