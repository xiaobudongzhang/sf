package dynamic

import "math"

func MaxSubArray(nums []int) int {
	sum := nums[0]
	res := 0
	for i:=0;i<len(nums);i++{
		if sum > 0 {
			sum +=nums[i]
		}else {
			sum = nums[i]
		}

		if sum > res {
			res = sum
		}
	}
	return res
}


func Rob(nums []int) int {
	last ,now := 0,0
	for i:=0;i<len(nums);i++ {
		last,now=now,  int(math.Max(float64(now), float64(last + nums[i])))
	}
	return now
}