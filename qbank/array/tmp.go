package array

import (
	"fmt"
	//"math"
	"sort"
)

func Intersect(nums1 []int, nums2 []int) []int {
	sum := make([]int, 0)
	position := 0
	for i := 0; i < len(nums1); i++ {

		for j := 0; j < len(nums2); j++ {
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

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func ThreeSum(nums []int) [][]int {

	var nums2 IntSlice
	numslen := len(nums)
	nums2 = make([]int, numslen)
	copy(nums2, nums)
	sort.Sort(nums2)

	var result [][]int
	var i int

	for i = 0; i < numslen; i++ {

		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
		low := i + 1
		hight := numslen - 1

		for low < hight {
			//fmt.Printf("low:%v-hight:%v\n",  low, hight)
			fmt.Printf("%v:%v:%v\n", nums[low], nums[hight], target)
			if nums[low]+nums[hight] == target {
				//fmt.Printf("%v:%v:%v\n", nums[low] , nums[hight],target)
				result = append(result, []int{nums[i], nums[low], nums[hight]})

				for low < hight && nums[low] == nums[low+1] {
					low++
				}

				for low < hight && nums[hight] == nums[hight-1] {
					hight--
				}
				low++
				hight--

			} else if nums[low]+nums[hight] < target {
				//fmt.Printf("low:%v-hight:%v\n",  low, hight)
				low++
			} else if nums[low]+nums[hight] > target {
				hight--
			} else {
				fmt.Printf("xyz\n")
			}
		}

	}

	return result
}
