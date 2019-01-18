package array


func Intersect(nums1 []int, nums2 []int) []int {
	sum := make([]int, 0)
	position := 0
	for i:=0;i<len(nums1);i++{

		for j:=0;j<len(nums2);j++{
				if j < position {
					continue
				}
			if nums1[i] == nums2[j] {
				position++
				sum = append(sum, nums1[i])
				break
			}
		}
	}
	return sum
}
