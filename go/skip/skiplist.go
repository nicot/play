package skip

type skipList struct {
	key   int
	value interface{}
	next  *skipList
	down  *skipList
}

// Skiplist is a skiplist
type Skiplist struct {
	underlying *skipList
}

// Set puts a key into the skiplist
func (s *Skiplist) Set(key int, value interface{}) {
	element := &skipList{
		key,
		value,
		nil,
		nil,
	}
	if s.underlying == nil {
		s.underlying = element
	}

}

// Get grabs a key from the skiplist
func (s *Skiplist) Get(key int) interface{} {
	v, present := get(s.underlying, key)
	if !present {
		panic("value not found")
	}
	return v
}

// Get grabs a key from the skiplist
func get(s *skipList, key int) (interface{}, bool) {
	if s == nil {
		return nil, false
	}
	if s.key == key {
		return key, true
	}
	if s.next == nil {
		return get(s.down, key)
	}
	if key > s.next.key {
		return get(s.next, key)
	}
	return get(s.down, key)
}
