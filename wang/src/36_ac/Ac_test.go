package _6_ac

import (
	"testing"
	"lib"
	"35_trie"
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
