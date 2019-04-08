package _7_greedy

import (
	//"fmt"
	"sort"
	"fmt"
)

type Node struct {
	l int
	r int
}
type Point []*Node

func (this Point)Len() int  {
	return len(this)
}

func (this Point)Less(i, j int) bool  {
	return this[i].l < this[j].l
}

func (this Point)Swap(i, j int)  {
	this[i], this[j] = this[j], this[i]
}
func Space(point Point)  {

	sort.Sort(point)

	start := 0
	end := 0

	//可选的进行比较，取最靠近左侧的区间
	for i := 0; i < len(point); i++ {

		if point[i].l > start && point[i].l < end {//说明区间已经被覆盖了
			continue
		}

		minv := point[i]
		end = minv.r

		//有没有比符合条件的区间更优的区间(核心...)
		for j:=i;j<len(point);j++ {
			if point[j].r < end {
				minv = point[j]
				i = j
				end = minv.r
			}
		}

		fmt.Printf("v:%v\n", minv)

	}



}
