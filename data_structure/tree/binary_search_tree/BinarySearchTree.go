package binary_search_tree

import (
	"container/list"
	"fmt"
)

//@todo 暂不支持重复树
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func newNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

func NewBinaryNode() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (tree *BinarySearchTree) Find(data interface{}) *Node {
	node := tree.root
	for node != nil {
		if data == node.data {
			return node
		}
		if data.(int) < node.data.(int) {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (tree *BinarySearchTree) Insert(data interface{}) bool {
	if tree.root == nil {
		tree.root = newNode(data)
		return true
	}
	if data == tree.Find(data) {
		return false
	}
	node := tree.root
	for node != nil {
		if data.(int) < node.data.(int) {
			if node.left == nil {
				node.left = newNode(data)
				return true
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = newNode(data)
				return true
			}
			node = node.right
		}
	}
	return true
}

func (tree *BinarySearchTree) Delete(data interface{}) bool {
	//找到要删除的节点
	p := tree.root
	pp := newNode(nil)
	pp = nil

	for p != nil {
		if data.(int) == p.data.(int) {
			break
		}
		pp = p
		if data.(int) < p.data.(int) {
			p = p.left
		} else {
			p = p.right
		}
	}

	if p == nil {
		return true
	}

	//如果有左子树与右子树,1.用右子树的最小节点替换要删除的值  2.删除掉被替换的节点
	if p.left != nil && p.right != nil {
		minP := p.right
		minPp := p
		//1.找到右子树的最小节点
		for minP.left != nil {
			minPp = minP
			minP = minP.left
		}
		//替换
		p.data = minP.data

		//2.删除掉被替换的节点
		p = minP
		pp = minPp
	}

	//如果有一个节点
	child := newNode(nil)
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {
		tree.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}

	return true
}

func (tree *BinarySearchTree) InOrderPrint() {
	if tree.root == nil {
		fmt.Println("empty tree")
		return
	}
	lists := list.New()
	lists.PushBack(tree.root)

	i := 0
	qlen := lists.Len()

	for lists.Len() > 0 {

		node := lists.Front()
		lists.Remove(node)
		nodeV := node.Value.(*Node)

		fmt.Printf(" %v ", nodeV.data)

		if nodeV.left != nil {
			lists.PushBack(nodeV.left)
		}
		if nodeV.right != nil {
			lists.PushBack(nodeV.right)
		}
		i++
		if qlen == i {
			qlen = lists.Len()
			i = 0
			fmt.Println("") //分层
		}

	}
	fmt.Println("") //分层

}
