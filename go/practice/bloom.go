package practice

const offset uint64 = 14695981039346656037
const prime uint64 = 1099511628211
const k = 5

type filter struct {
	m uint64
	k uint64
	b []byte
}

// simple 64 bit FNV hash
func hash(input []byte) uint64 {
	hash := offset
	for _, e := range input {
		hash ^= uint64(e)
		hash *= prime
	}
	return hash
}

// Generate the nth hash of the input
func hashi(input []byte, n uint64) uint64 {
	h := hash(input)
	h1 := h >> 32         // first half of hash
	h2 := h & (1<<32 - 1) // second half of hash
	return uint64(h1 + h2*n)
}

// add elem to the filter
func (f filter) add(elem []byte) {
	var i uint64
	for i = 0; i < f.k; i++ {
		v := hashi(elem, i) % f.m
		f.b[v] = 1
	}
	return
}

// test whether elem is in the filter
func (f filter) test(elem []byte) bool {
	var i uint64
	for i = 0; i < f.k; i++ {
		v := hashi(elem, i) % f.m
		if f.b[v] == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	var f filter
// 	f.m = 1024
// 	f.b = make([]byte, f.m)
// 	f.k = 5

// 	fmt.Println(f.test([]byte("a")))
// 	f.add([]byte("a"))
// 	fmt.Println(f.test([]byte("a")))
// }
