package n

import (
	"github.com/xiaobudongzhang/sf/algorithm/sort/nlogn"
	"github.com/xiaobudongzhang/sf/lib"
)

func BucketSort(a []int, bucketSize int) {
	if len(a) < 1 {
		return
	}
	min := lib.Min(a)
	max := lib.Max(a)
	if bucketSize < 2 {
		bucketSize = 2
	}
	//分桶
	k := (max - min) / bucketSize
	buckets := make(map[int][]int)
	for _, item := range a {
		index := (item - min) / k
		buckets[index] = append(buckets[index], item)
	}
	//排序
	i := 0
	for j := 0; j < k; j++ { //不能用range
		nlogn.QuickSort(buckets[j], len(buckets[j]))
		//合并
		for _, item := range buckets[j] {
			a[i] = item
			i++
		}
	}

}
