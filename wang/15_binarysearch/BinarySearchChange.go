package _5_binarysearch


func FirstEqual(a []int) int {
	lo := 0
	hi := len(a) -1
	s := 8
	for lo <= hi{
		mid := lo + ((hi - lo) / 2)
		if a[mid] == s {
			if ((mid == 0) || a[mid -1] != s){//前一个值不等于说明现在的是第一个
				return mid
			}else{
				hi = mid -1
			}
		}else if(a[mid] < s){
			lo = mid + 1
		}else {
			hi = mid - 1
		}
	}

	return -1
}

func LastQqual()  {
	
}

func FirstGtOrEqual()  {
	
}

func LastLessOrEqual()  {
	
}
