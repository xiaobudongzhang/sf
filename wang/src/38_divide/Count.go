package _8_divide

import "fmt"

var num int
func Count(arr[]int) int {
	num = 0
	mergeSortCounting(arr, 0, len(arr) -1)
	return num
}

func mergeSortCounting(arr []int, p int, r int)  {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSortCounting(arr, p, q)
	mergeSortCounting(arr, q+1, r)
	merge(arr, p, q, r)
}

func merge(arr []int, p int, q int, r int)  {
	i,j,k := p, q+1,0
	tmpArr := make([]int, r-p+1)

	for i <= q && j <= r  {
		if arr[i] <= arr[j]{
			tmpArr[k] = arr[i]
			k++
			i++
		}else {
			num += (q-i+1)//从i到结尾肯定都大于a[j]
			tmpArr[k] = arr[j]
			k++
			i++
		}
	}

	for i<=q  {
		tmpArr[k] = arr[i]
		k++
		i++
	}
	for j<=r {
		tmpArr[k] = arr[j]
		k++
		j++
	}

	copy(arr, tmpArr)
}
