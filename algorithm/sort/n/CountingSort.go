package n

import (
	"github.com/xiaobudongzhang/sf/lib"
)

func CountingSort(a []int) {
	if len(a) < 2 {
		return
	}
	min := lib.Min(a)
	max := lib.Max(a)

	c := make([]int, max-min+1) //存储数字的数量
	for _, item := range a {
		c[item-min]++
	}
	//顺序求和
	sum := 0
	for k, item := range c {
		sum += item
		c[k] = sum
	}
	r := make([]int, len(a)) //排序的数字

	for i := len(a) - 1; i >= 0; i-- {
		c[a[i]-min]--
		r[c[a[i]-min]] = a[i]
	}
	copy(a, r)
}
