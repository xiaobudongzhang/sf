package _5_binarysearch

import (
	"math"
	//"fmt"
)

func BinarySearch(a []int, search int) int {
	if len(a) < 1 {
		return -1
	}
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := (lo + hi) / 2
		if search == a[mid] {
			return mid
		} else if search < a[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func Sqr(a float64) float64 {
	result := make([]float64, 1)
	sqrChild(a, 1.0, a/2, result)
	return result[0]
}

func sqrChild(a float64, lo float64, hi float64, result []float64) {
	mid := lo + ((hi - lo) / 2)
	if math.Abs(a-mid*mid) < 0.0000001 {
		result[0] = mid
		return
	} else if mid*mid < a {
		sqrChild(a, mid, hi, result)
	} else {
		sqrChild(a, lo, mid, result)
	}
}
