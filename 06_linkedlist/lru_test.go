package _6_linkedlist

import (
	"testing"
)

func TestLRU_Find(t *testing.T) {
	PrintFunc()
	lru := NewLRU(4)

	lru.Find("a")
	lru.Find("b")
	lru.Find("c")
	lru.Find("d")
	lru.linkedList.Print()
	lru.Find("c")
	lru.linkedList.Print()

	lru.Find("a")
	lru.linkedList.Print()

	lru.Find("d")
	lru.linkedList.Print()
}
