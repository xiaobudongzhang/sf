package _6_ac

import (
	"35_trie"
	"lib"
	"testing"
)

func TestAc_Macth(t *testing.T) {
	lib.PrintFunc()

	trie := _5_trie.NewTrie()
	trie.Insert("abcd")
	trie.Insert("bcd")
	trie.Insert("c")

	trie.BuildFailurePointer()

	//trie.InOrder()
	trie.Match("xbcdyax")
}
