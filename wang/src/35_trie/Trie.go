package _5_trie

import (
	"fmt"
	"09_queue"
)

type TrieNode struct {
	data byte
	children [27]*TrieNode//为了区分初始化的byte0 与a的区别 多加一位，默认是数组0 a是1
	isEndingChar bool
	fail *TrieNode
	length int
}

type Trie struct {
	root *TrieNode
	startChar byte
}

func NewTrieNode(data byte) *TrieNode {
	return &TrieNode{
		data:data,
		isEndingChar:false,
		length:-1,
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
	p.length = len(text)
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
//按层级构造失败指针
func (this *Trie) BuildFailurePointer(){
	p := this.root
	queue := _9_queue.NewLinkedListQueue()
	this.root.fail = nil

	queue.EnQueue(p)
	for !queue.IsEmpty(){
		p := queue.DeQueue().(*TrieNode)

		for _,v := range p.children  {
			if v == nil || v.data == 0 {
				continue
			}
			if p == this.root {//如果是根节点的话，则子类的fail都指向根节点
				v.fail = this.root
				queue.EnQueue(v)
				continue
			}
			q := p.fail
			for q != nil  {//核心点
				index := v.data - (this.startChar)
				qc := q.children[index]

				if qc != nil {//两者相等 pc 跟qc相等
					v.fail = qc
					break
				}

				q = q.fail
			}
			if q == nil {//如果没有失败的子节点了就指向根节点
				v.fail = this.root
			}
			queue.EnQueue(v)
		}
	}
}

func (this *Trie) Match(a string)  {

	p := this.root

	for i,v := range a{
		index := v - int32(this.startChar)
		for p.children[index] == nil && p != this.root {
			p = p.fail //失败的话查看最长子串有没有匹配的情况
		}
		p = p.children[index]
		if p == nil {//如果没有匹配的从root重新开始,并且查找字符向下走
			p = this.root
			continue
		}
		tmp := p
		for tmp != this.root  {//整个for循环结束，查找字符往下走
			if tmp.isEndingChar == true {
				pos := i - tmp.length + 1
				fmt.Printf("匹配的起始下标：%v;长度：%v",pos, tmp.length)
			}
			tmp = tmp.fail
		}


	}
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
		if p.fail != nil{
			fmt.Printf("(%+v)  ", string(p.fail.data))
		}
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