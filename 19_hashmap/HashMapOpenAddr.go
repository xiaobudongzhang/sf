package _9_hashmap

import (
	"lib"
	"strconv"
	"fmt"
)

type HashMap struct {
	used int
	datas []int
	size int
}
//暂不考虑扩容
func NewHashTable(size int) *HashMap {
	return &HashMap{
		used:0,
		datas:make([]int, size, size),
		size:size,
	}
}

func (this *HashMap)Insert(v int)  bool{
	index := lib.Hash33(strconv.Itoa(v), this.size)
	fmt.Printf("index:%v\n", index)
	i := index
	for this.datas[i]  > 0 {
		i = (i + 1) % this.size
		if i == index {//说明循环一圈了
			return false
		}
	}
	this.used++
	this.datas[i] = v
	return true
}

func (this *HashMap)Delete(v int) int{

	index := lib.Hash33(strconv.Itoa(v), this.size)
	i := index

	for this.datas[i]  != v {
		i = (i + 1) % this.size
		if i == index {//说明循环一圈了
			return -2
		}
	}
	this.used--
	this.datas[i] = -1

	return i
}

func (this *HashMap)Find(v int) int {
	index := lib.Hash33(strconv.Itoa(v), this.size)
	i := index

	for this.datas[i]  != v && this.datas[i] != 0{
		i = (i + 1) % this.size
		if i == index {//说明循环一圈了
			return -2
		}
	}

	return i
}
func (this *HashMap)Print()  {
	fmt.Printf("used:%v,datas:%v", this.used,this.datas, )
}