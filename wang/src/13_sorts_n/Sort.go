package _3_sorts_n

import (
	//"fmt"
	 "math"
)

//import "fmt"
//@todo 如果桶太大继续分桶，取总数量作为桶的下标，这样就能解决继续分桶时下标重复的问题，
func BucketSort(a []int)  {
	min := min(a)
	max := max(a)
	k := (max - min) / 10

	bucketMap := make(map[int][]int)
	for i := 0; i < len(a) ; i++ {
		index := a[i] / k
		tmpArr := bucketMap[index]
		tmpArr = append(tmpArr, a[i])
		bucketMap[index] = tmpArr
	}

	for i:= 0; i < max /k ;i++{
		bubbleSort(bucketMap[i], len(bucketMap[i]))
	}
	z := 0
	for i:= 0; i < max /k + 1;i++{
		for j := 0;j < len(bucketMap[i]); j++{
			a[z] = bucketMap[i][j]
			z++
		}
	}
}

func CountingSort(a []int)  {
	if len(a) < 2{
		return
	}
	min := min(a)
	len2 := max(a) - min + 1
	countArr := make([]int, len2)
	//计数
	for i := 0; i < len(a); i++  {
		index := a[i] - min
		countArr[index]++
	}
	//求和
	sum := 0
	for i := 0; i < len(countArr); i++{
		sum += countArr[i]
		countArr[i] = sum
	}
	//放回原数组
	newA := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		countIndex := a[i] - min
		index := countArr[countIndex]
		newA[index-1] = a[i]
		countArr[countIndex]--
	}
	copy(a, newA)
}
//基数排序
func RadixSort(a []int)  {
	maxLoop := 3
	for i := 0; i < maxLoop; i++ {
		pval := math.Pow10(i)

		//从最后一位进行计数排序
		countArr := make([]int, 10)
		//计数
		for i := 0; i < len(a); i++  {
			index := (a[i] / int(pval)) % 10
			countArr[index]++
		}

		//求和
		for i := 1; i < 10; i++{
			countArr[i] += countArr[i-1]
		}
		//排序
		newA := make([]int, len(a))
		for i := len(a) - 1; i >= 0; i-- {
			countIndex := (a[i] / int(pval)) % 10
			index := countArr[countIndex]
			newA[index-1] = a[i]
			countArr[countIndex]--
		}
		//放回原数组
		copy(a, newA)
	}
}

func bubbleSort(a []int, n int)  {
	for i := 0; i < n; i++  {
		change := false
		for j := 0; j < n - i -1; j++  {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
				change = true
			}
		}
		if !change {
			break
		}
	}
}

func min(a []int)  int{
	min := a[0]
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}
	return min
}

func max(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
