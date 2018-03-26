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

}

// Get grabs a key from the skiplist
func (s *Skiplist) Get(key int) interface{} {
	return nil
}
