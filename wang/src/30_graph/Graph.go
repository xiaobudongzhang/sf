package _0_graph

import "fmt"

type Node struct {
	v interface{}
	next *Node
}

type Graph struct {
	datas []*Node
}

func NewGraph(size int) *Graph {
	datasArr := make([]*Node, size,size)
	for i:=0;i<size ;i++  {
		datasArr[i] = &Node{
			v:nil,
			next:nil,
		}
	}
	//第一个node存储的是a
	return &Graph{
		datasArr,
	}
}

func NewNode(v interface{}) *Node {
	return &Node{
		v:v,
		next:nil,
	}
}

func (this *Graph) InsertA(a interface{}){

	for i:=0 ;i<len(this.datas);i++{
			if this.datas[i].v == nil{
				continue
			}
			if this.datas[i].v == a {
				return
			}
	}

	for i:=0 ;i<len(this.datas);i++{
		if this.datas[i].v == nil {
			this.datas[i].v = a
			return
		}
	}
}

func (this *Graph) InsertOfA(a interface{}, bofa interface{}){

	index := -1
	for i:=0 ;i<len(this.datas);i++{
		if this.datas[i].v == a {
			index = i
			break
		}
	}
	if index < 0 {
		return
	}
	pNode := this.datas[index]
	for pNode.next != nil {
		pNode = pNode.next
	}
	newNode := NewNode(bofa)
	pNode.next = newNode
}

func (this *Graph) Print() {

	for i:=0;i<len(this.datas) ;i++  {

		if this.datas[i].v == nil{
			continue
		}
		if this.datas[i].v.(int) < 1{
			continue
		}
		fmt.Printf("%v:",this.datas[i].v)
		pNode := this.datas[i]
		for pNode != nil {
			fmt.Printf("   %v ", pNode.v)
			pNode = pNode.next
		}
		fmt.Printf("\n")
	}
}
