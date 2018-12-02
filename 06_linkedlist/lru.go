package _6_linkedlist

//import "fmt"

/*
1. 如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。

2. 如果此数据没有在缓存链表中，又可以分为两种情况：
如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部。
如果此时缓存未满，则将此结点直接插入到链表的头部；
*/

type LRU struct {
	linkedList *LinkedList
	length uint
}

func NewLRU(length uint)  *LRU{
	return &LRU{NewLinkedList(), length}
}

func (this *LRU) Find(value interface{}) interface{} {
	pNode := this.linkedList.FindByVal(value)

	// in cache
	if pNode != nil {
		if pNode == this.linkedList.head {
			return value
		}
		this.linkedList.DeleteNode(pNode)
		this.linkedList.InsertToHead(value)
		return value
	}
	//not in cache
	if this.linkedList.length >= this.length {
		// cache full
		this.linkedList.DeleteLastNode()
		this.linkedList.InsertToHead(value)
		return value
	}
	//cache not full
	this.linkedList.InsertToTail(value)
	this.length++
	return value
}