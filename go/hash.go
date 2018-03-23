package main

import "fmt"

func main() {
	fmt.Println("hello!")
}

type linkedList struct {
	key   int
	value interface{}
	next  *linkedList
}

// Table is a hash table.
type Table struct {
	elements []*linkedList
	values   int
	slots    int
}

const loadFactor = 0.9

// Put adds an element to the hash table.
func (t *Table) Put(k int, v interface{}) {
	element := &linkedList{
		k,
		v,
		nil,
	}

	if len(t.elements) == 0 {
		t.values++
		t.slots++
		t.elements = append(t.elements, element)
		return
	}

	if float64(t.values)/float64(t.slots) > loadFactor {
		t.grow()
	}

	put(t.elements, k, t.hash(k), v)
	t.values++
}

func (t *Table) hash(n int) int {
	// TODO use fnv or something to randomize key.
	return n % t.slots
}

func put(elements []*linkedList, k int, hash int, v interface{}) {
	element := &linkedList{
		k,
		v,
		nil,
	}

	ll := elements[hash]
	if ll == nil {
		elements[hash] = element
		return
	}

	for ll.next != nil {
		ll = ll.next
	}
	ll.next = element
}

const growthFactor = 2

func (t *Table) grow() {
	next := make([]*linkedList, t.slots*2)
	t.slots = t.slots * 2
	for _, node := range t.elements {
		for node != nil {
			put(next, node.key, t.hash(node.key), node.value)
			node = node.next
		}
	}
	t.elements = next
}

// Get a key from the hash table.
func (t *Table) Get(k int) interface{} {
	hash := t.hash(k)

	if hash > len(t.elements) {
		panic("hash out of range")
	}

	ll := t.elements[hash]
	for ll != nil {
		if ll.key == k {
			return ll.value
		}
		ll = ll.next
	}
	return nil
}

// Delete something from the table.
func (t *Table) Delete(k int) {
	hash := t.hash(k)
	if hash > len(t.elements) {
		panic("hash out of range")
	}

	ll := t.elements[hash]
	if ll.key == k {
		t.elements[hash] = nil
		return
	}
	for ll.next != nil {
		if ll.next.key == k {
			// it is OK if ll.next.next is nil, and ll.next is not nil.
			ll.next = ll.next.next
			return
		}

		ll = ll.next
	}
	panic("tried to delete an element that doesn't exist")
}
