package _2_sorts_nlongn

import "fmt"

func FindLastK(a []int, k int) int {
	if k > len(a) {
		return -1
	}
	kval := make([]int, 1)
	findLastKChild(a, 0, len(a)-1, k, kval)
	return kval[0]
}

//求解的结果放到传引用里
func findLastKChild(a []int, p int, r int, k int, kval []int) {
	q := partition2(a, p, r)
	if q+1 == k {
		kval[0] = a[q]
		return
	}
	if k > q+1 {
		findLastKChild(a, q+1, r, k, kval)
	} else {
		findLastKChild(a, p, q-1, k, kval)
	}
	return
}

func partition2(a []int, p int, r int) int {
	provit := a[r]
	i := p
	for j := p; j < r; j++ {
		if a[j] > provit {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	fmt.Printf("a[i]:%v", a[i])
	a[i], a[r] = a[r], a[i]
	fmt.Printf("provit:%v- i:%v\n", provit, i)
	return i
}
