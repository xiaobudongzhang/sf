package n

import (
	"github.com/xiaobudongzhang/sf/lib"
)

//基数排序
func RadixSort(a []int) {
	max := lib.Max(a)
	//min := lib.Min(a)

	maxLoop := max
	k := 1

	for ; maxLoop >= 1; maxLoop = maxLoop / 10 { //按照数字的位数循环
		//进行基数排序
		//计数
		c := make([]int, 10)
		for _, item := range a {
			index := (item / k) % 10
			c[index]++
		}
		//顺序求和
		for j := 1; j < len(c); j++ {
			c[j] = c[j-1] + c[j]
		}
		//排序
		r := make([]int, len(a))
		for i := len(a) - 1; i >= 0; i-- {
			index := (a[i] / k) % 10
			c[index]--
			r[c[index]] = a[i]
		}
		copy(a, r)
		k = k * 10
	}
}
