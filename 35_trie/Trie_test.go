package _5_trie

import (
	"testing"
	"lib"
	//"fmt"
	"fmt"
)

func TestTrie_Insert(t *testing.T) {
	lib.PrintFunc()

	trie := NewTrie()
	trie.Insert("how")
	trie.Insert("hi")
	trie.Insert("her")
	trie.Insert("hello")
	trie.Insert("so")
	trie.Insert("see")
	trie.Insert("alas")

	trie.InOrder()


	f1 := trie.Find("see")
	f2 := trie.Find("se")
	fmt.Printf("f1:%v,f2:%v\n",f1, f2)
}
