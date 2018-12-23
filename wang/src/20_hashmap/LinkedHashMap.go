package _0_hashmap

import (
	"fmt"
)

type Node struct {
	prev *Node
	key interface{}
	val interface{}
	next *Node
	hnext *Node
}

type HashMap struct {
	htables []*Node
	size int
	hhead *Node
	htail *Node
}

func newNode(key interface{}, val interface{}) *Node {
	return &Node{
		key:key,
		val:val,
	}
}

func NewHashMap(size int) *HashMap {
	nodeList := make([]*Node, size, size)
	for i:=0;i<size ; i++ {
		nodeList[i] =  &Node{}//方便初始化散列表的修改
	}
	return &HashMap{
		size:size,
		htables:nodeList,
	}
}
func (this *HashMap) Insert(k interface{}, v interface{})  {
	khash := this.Hash(k)
	hlist := this.htables[khash]
	cur := hlist
	for cur != nil {
		if cur.key == k{
			break
		}
		if cur.hnext == nil{
			break
		}
		cur = cur.hnext
	}

	if cur.key == k {//如果找到了 1.散列表修改内容 2.链表放到表尾部
		//更新hhead
		if this.hhead.next != nil && this.hhead.key == k{
			this.hhead = this.hhead.next
		}

		cur.val = v //1

		//2
		this.htail.next = cur
		cur.prev = this.htail
		cur.next = nil
		this.htail = cur

		return
	}

	//没找到加到1.散列加到尾部2.链表加到尾部
	nNode := newNode(k, v)
	//1
	cur.hnext = nNode
	//2.
	if this.htail != nil{
		this.htail.next = nNode
	}
	nNode.prev = this.htail
	this.htail = nNode
	if this.hhead == nil{
		this.hhead = nNode
	}
}

func (this *HashMap) Find(k interface{}) interface{}{
	khash := this.Hash(k)
	hlist := this.htables[khash]
	cur := hlist
	for cur != nil {
		if cur.key == k{
			break
		}
		if cur.hnext == nil{
			break
		}
		cur = cur.hnext
	}
	return cur.val
}

func (this *HashMap) Delete(k interface{}) {
	khash := this.Hash(k)
	hlist := this.htables[khash]
	cur := hlist.hnext
	prev := hlist
	for cur != nil {
		if cur.key == k{
			break
		}
		prev = prev.hnext
		cur = cur.hnext
	}
	//未找到
	if cur == nil{
		return
	}
	if cur.key != k {
		return
	}
	fmt.Printf("prev.key:%v\n",prev.key)
	//1.删除散列表 2.删除链表
	prev.hnext = prev.hnext.hnext//1
	//2
	//如果删除的是尾节点
	if cur == this.htail {
		this.htail = cur.prev
	}
	//如果删除的是头节点
	if cur == this.hhead{
		this.hhead = cur.next
	}
	//如果删除的节点没有前驱节点，如果删除的节点没有后驱节点
	if cur.prev != nil{
		cur.prev.next = cur.next
	}
	if cur.next != nil{
		cur.next.prev = cur.prev
	}
}

func (this *HashMap) Hash(k interface{}) int{
	//time33
	hash := 5381
	for _,v := range k.(string) {
		hash += (hash << 5) + int(v)
	}
	return (hash & 0x7FFFFFFF) % this.size
}

func (this *HashMap) Print() {
	fmt.Printf("\nhashtable:\n")
	for i:=0; i<this.size; i++{
		fmt.Printf("---i:%v", i)
		node := this.htables[i]
		for node != nil {
			fmt.Printf("(%v:%v)", node.key, node.val)
			node = node.hnext
		}
		fmt.Printf("---\n")
	}
	fmt.Printf("\nlinked:\n")
	head := this.hhead
	fmt.Printf("---")
	for  head!=nil {
		fmt.Printf("(%v:%v)", head.key, head.val)
		head = head.next
	}
	fmt.Printf("---\n")
}