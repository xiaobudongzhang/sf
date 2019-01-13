package array



func RemoveDuplicates(nums []int) int {
	//将相同的元素都移动到后面
	len := len(nums)

	i := 1
	for i < len {
		if nums[i] == nums[i-1] {//将i位的元素删除并且将所有元素后移，最后将i位元素放在最后一位或者不处理
			if nums[i] == nums[len-1] {
				break
			}
			for j:=i;j<len-1;j++{
				nums[j] = nums[j+1]
			}
		}else{
			i++
		}
	}
	return i
}