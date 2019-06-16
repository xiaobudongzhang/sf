package trie_tree

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestTrie_Find(t *testing.T) {
	lib.PrintFunc()
	tree := NewTrie()

	tree.Insert("he")
	tree.Insert("her")
	tree.Insert("his")

	tree.Print()
	fmt.Printf("he:%v,she:%v", tree.Find("he"), tree.Find("she"))
}
