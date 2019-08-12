package sort

var topK int
var k int
func FindTopK(a []int, n int,fk int) int  {
	topK = -1
	if fk > n {
		return topK
	}
	k = fk

	findTopKChild(a, 0, n -1)
	return topK
}

func findTopKChild(a []int, p int, r int)  {
	if p>=r {
		return
	}
	q := partition(a, p, r)

	if q + 1 == k {//k+1 因为下标从0开始
		topK = a[q]
		return
	}


	if q + 1 < k {
		findTopKChild(a, q + 1, r)
	} else {
		findTopKChild(a, p, q -1)
	}
}

func partition(a []int, p int, r int) int  {
	i := p
	privot := a[r]
	for j := p;j<r ;j++  {
		if a[j] > privot {
			a[j],a[i] = a[i],a[j]
			i++
		}
	}

	a[i],a[r] = a[r],a[i]
	return i
}