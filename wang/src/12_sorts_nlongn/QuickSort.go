package _2_sorts_nlongn

//import "fmt"

//1.quick_sort(p...r) = quick_sort(p...q-1)+quick_sort(q+1,r)
//2.p>=r
func QuickSort(a []int)  {
	if len(a) < 2 {
		return
	}
	quickSortChild(a, 0, len(a) - 1)
}

func quickSortChild(a []int, p int, r int)  {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	quickSortChild(a, p, q - 1)
	quickSortChild(a, q + 1, r)
}

func partition(a []int, p int, r int) int {
	pivot := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}
