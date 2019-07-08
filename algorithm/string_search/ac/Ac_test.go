package ac

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestSearch(t *testing.T) {
	lib.PrintFunc()


	trie := NewTrie()
	trie.Insert("how")
	trie.Insert("hi")
	trie.Insert("her")
	trie.Insert("hello")
	trie.Insert("so")
	trie.Insert("see")
	trie.Insert("alas")


	trie.BuildFailurePointer()
	trie.Search("helloiseeyouhealas")
}
