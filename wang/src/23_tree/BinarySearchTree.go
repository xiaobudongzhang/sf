package _3_tree

import "fmt"

type BST struct {
	*BinaryTree
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(root interface{}, compareFun func(v, node interface{}) int) *BST {
	return &BST{
		BinaryTree:NewBinaryTree(root),
		compareFunc:compareFun,
	}
}

func (this *BST) Find(v interface{}) *Node {
	pNode := this.root
	for pNode != nil {
		res := this.compareFunc(v, pNode.data)
		if res > 0 {
			pNode = pNode.right
		}else if (res < 0){
			pNode = pNode.left
		}else {
			return pNode
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	pNode := this.root
	for pNode != nil {
		res := this.compareFunc(v, pNode.data)
		if res == 0 {
			return false
		}
		//right
		if res > 0 {
			if pNode.right == nil {
				pNode.right = NewNode(v)
				break
			}
			pNode = pNode.right
		}else {
			//left
			if pNode.left == nil{
				pNode.left = NewNode(v)
				break
			}
			pNode = pNode.left
		}
	}
	return true
}


func (this *BST) Delete(v interface{}) bool {
	//判断删除的节点的情况（1.没有子节点 2.有一个子节 3. 有两个子节点）
	pNode := this.root
	ppNode := this.root//父节点
	for pNode != nil && this.compareFunc(v, pNode.data) != 0{
		ppNode = pNode
		if this.compareFunc(v, pNode.data) > 0 {
			pNode = pNode.right
		}else {
			pNode = pNode.left
		}
	}
	if pNode == nil {
		return false
	}
	//1.没有子节点
	if pNode.left == nil && pNode.right == nil {
		if ppNode.right == pNode {
			ppNode.right = nil
		}else {
			ppNode.left = nil
		}
		return true
	}
	//2.有一个子节点
	if pNode.left != nil && pNode.right == nil {
		if ppNode.right == pNode {
			ppNode.right = pNode.left
		}else {
			ppNode.left = pNode.left
		}
		return true
	}

	if pNode.right != nil && pNode.left == nil {
		if ppNode.right == pNode {
			ppNode.right = pNode.right
		}else {
			ppNode.left = pNode.right
		}
		return true
	}


	//3.有两个子节点（找到右子树的最小节点，替换掉要删除的节点）
	minNode := pNode.right
	minNodePP := pNode
	for minNode.left != nil {
		minNodePP = minNode
		minNode = minNode.left
	}
	//替换
	pNode.data = minNode.data
	//删除最小节点
	if minNodePP.left == minNode{
		minNodePP.left = nil
	}else {
		minNodePP.right = nil
	}


	return true
}

func (this *BST) Print(root *Node)  {
	fmt.Printf("%v\n", root.data)
	if root.left != nil {
		this.Print(root.left)
	}

	if root.right != nil {
		this.Print(root.right)
	}
}