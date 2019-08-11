package nlogn

func MergeSort(a []int, n int)  {
	if n <= 1 {
		return
	}
	mergeSortChild(a, 0, n -1)
}

func mergeSortChild(a []int, p int, r int)  {
	if p >= r {
		return
	}
	q := (p+r)/2
	mergeSortChild(a, p, q)
	mergeSortChild(a, q + 1, r)
	merge(a, p, q, r)
	return
}

func merge(a []int, p int, q int, r int)  {
	tmp := make([]int, r - p + 1) //创建临时数组
	i := p
	j := q + 1
	k := 0

	for i <= q && j <= r {
		if a[i] < a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	for ;i<=q ; i++ {
		tmp[k] = a[i]
		k++
	}

	for ;j<=r ;j++  {
		tmp[k] = a[j]
		k++
	}
	//将排序后的数组复制到元数组
	m := 0
	for l := p; l <= r; l++  {
		a[l] = tmp[m]
		m++
	}
}
