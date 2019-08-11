package nlogn

func QuickSort(a []int, n int)  {
	if n <= 1 {
		return
	}
	quickSortChild(a, 0, n -1)
}

func quickSortChild(a []int, p int, r int) {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	//分成前后两部分
	quickSortChild(a, p, q -1)//前半部分是q-1
	quickSortChild(a, q + 1, r)

}
//类似插入排序
func partition(a []int, p int, r int) int {
	i := p//已排序边界
	pivot := a[r]//取最后一个元素做分界点
	for j := p; j < r ;j++  {
		if a[j] < pivot {//添加到排序的分区里
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i] //交换中间分界点
	return i
}