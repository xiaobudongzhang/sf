package trie_tree

import (
	"container/list"
	"fmt"
)

type  TrieNode struct {
	data byte
	children [26]*TrieNode
	isEndingChar bool //是否是完整的字符串,比如插入he,her 则e也是字符结尾
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root:newTrieNode('/'),
	}
}

func newTrieNode(data byte) *TrieNode  {
	return &TrieNode{
		data:data,
	}
}

func (tree *Trie) Insert(text string) bool{
	node := tree.root
	for _, v :=range text {
		index := v - 'a'
		if node.children[index] == nil {//子类不存在则创建
			node.children[index] = newTrieNode(byte(v))
		}
		node = node.children[index]//走到下一节点
	}
	node.isEndingChar = true
	return true
}

func (tree *Trie) Find(text string) bool {
	node := tree.root
	for _, v :=range text {
		index := v - 'a'
		if node.children[index] == nil {//子类不存在则创建
			return false
		}
		node = node.children[index]//走到下一节点
	}
	return node.isEndingChar
}

func (tree *Trie) Print()  {
	lists := list.New()
	p := tree.root
	lists.PushBack(p)
	len := 1
	i := 0

	for lists.Len() > 0 {
		node := lists.Front()
		lists.Remove(node)
		nodev := node.Value.(*TrieNode)
		i++

		fmt.Printf(" %v ", string(nodev.data))
		for _, child  := range nodev.children  {
			if child != nil {
				lists.PushBack(child)
			}
		}
		if len == i {
			fmt.Println("")
			len = lists.Len()
			i = 0
		}
	}

}