package _2_sorts_nlongn

//import "fmt"

//1.merge_sort(p..r) = merge(merge_sort(p...q), merge(q+1...r))
//2.p>=r
func MergeSort(a []int) {
	if len(a) < 1 {
		return
	}
	mergeSortChild(a, 0, len(a)-1)
}

func mergeSortChild(a []int, p int, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mergeSortChild(a, p, q)
	mergeSortChild(a, q+1, r)
	merge(a, p, q, r)
}

func merge(a []int, p int, q int, r int) {
	newArr := make([]int, len(a))
	pi := p
	pj := q + 1
	k := 0
	for pi <= q && pj <= r {
		if a[pi] > a[pj] {
			newArr[k] = a[pj]
			pj++
		} else {
			newArr[k] = a[pi]
			pi++
		}
		k++
	}
	start := pi
	end := q
	if pj <= r {
		start = pj
		end = r
	}
	for ; start <= end; start++ {
		newArr[k] = a[start]
		k++
	}

	copy(a[p:r+1], newArr)
}
