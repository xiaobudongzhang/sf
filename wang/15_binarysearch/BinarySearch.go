package _5_binarysearch

func BinarySearch(a []int, search int) int {
	if len(a) < 1 {
		return -1
	}
	lo := 0
	hi := len(a) -1
	for lo <= hi{
		mid := (lo + hi) / 2
		if (search == a[mid]) {
			return mid
		} else if(search < a[mid]){
			hi = mid - 1
		}else {
			lo = mid + 1
		}
	}
	return -1
}
