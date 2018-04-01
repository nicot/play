package main

import (
	"fmt"
	"io"
)

type Trie struct {
	elements map[byte]*Trie
	count int
}

func main() {
	var s string
	fmt.Scanln(&s)
	trie := Trie{
		make(map[byte] *Trie),
		0,
	}

	for {
		var operation, name string
		_, err := fmt.Scanf("%s %s", &operation, &name)
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		if operation == "add" {
			trie.add(name)
		} else if operation == "find" {
			trie.find(name)
		} else {
			panic(operation)
		}
	}
}

func (trie *Trie) add(name string) {
	trie.count++

	if name == "" {
		return
	}

	nextName := name[1:]
	c := name[0]

	nextNode := trie.elements[c]
	if nextNode == nil {
		nextNode = &Trie{
			make(map[byte] *Trie),
			0,
		}
		trie.elements[c] = nextNode
	}

	nextNode.add(nextName)
}

func (trie *Trie) find(name string) {
	if trie == nil {
		fmt.Println(0)
		return
	}
	if name == "" {
		fmt.Println(trie.count)
		return
	}

	nextNode := trie.elements[name[0]]
	nextName := name[1:]

	nextNode.find(nextName)
}