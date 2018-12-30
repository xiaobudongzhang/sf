package _5_trie

import (
	"fmt"
	"09_queue"
)

type TrieNode struct {
	data byte
	children [27]*TrieNode//为了区分初始化的byte0 与a的区别 多加一位，默认是数组0 a是1
	isEndingChar bool
}

type Trie struct {
	root *TrieNode
	startChar byte
}

func NewTrieNode(data byte) *TrieNode {
	return &TrieNode{
		data:data,
		isEndingChar:false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		NewTrieNode('/'),
		'`',
	}
}

func (this *Trie)Insert(text string)  {
	p := this.root

	for _, v := range text {

		index := v - int32(this.startChar)
		if p.children[index] == nil {
			p.children[index] = NewTrieNode(byte(v))
		}

		p = p.children[index]
	}
	//完整字符串的前缀子串都存默认 false,这样查找的时候，就可以判断是否是完整的子串了
	p.isEndingChar = true
}

func (this *Trie)Find(text string)  bool{
	p := this.root

	for _, v := range text {

		index := v - int32(this.startChar)
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}

	if p.isEndingChar == false {
		return false
	}
	return true
}


//按层遍历

func (this *Trie)InOrder() int{
	p := this.root
	q := _9_queue.NewLinkedListQueue()

	q.EnQueue(p)

	qlen := q.Len()
	i := 0
	height := 0
	fmt.Printf("level:0--")
	for !q.IsEmpty(){
		p := q.DeQueue().(*TrieNode)
		fmt.Printf("%+v ", string(p.data))

		for _,v := range p.children  {

			if v == nil || v.data == 0 {
				continue
			}

			q.EnQueue(v)
		}

		i++
		if i == qlen {
			height++
			fmt.Printf("\n")
			qlen = q.Len()
			if qlen > 0 {
				fmt.Printf("level:%v--", height)
			}
			i = 0
		}
	}
	fmt.Printf("height:%v\n", height)
	return height
}