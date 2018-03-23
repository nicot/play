package main

// Skiplist is a skiplist
type Skiplist struct{}

// Comparable used to compare keys for the skiplist.
type Comparable interface {
	// Compare should return -1 if the receiver is less than the parameter,
	// 0 if they are the same and 1 if the receiver is greater than the parameter.
	Compare(k Comparable) int
}

// Get pulls a key from the skiplist
func (s *Skiplist) Get(key Comparable) {

}

// Set puts a key into the skiplist
func (s *Skiplist) Set(key Comparable, value interface{}) {
	key.Compare(key)

}
