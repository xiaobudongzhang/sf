package array

func SingleNumber(nums []int) int {
	//异或处理
	result := 0
	for  i:=0;i <len(nums);i++{
		result ^= nums[i]
	}
	return result
}
