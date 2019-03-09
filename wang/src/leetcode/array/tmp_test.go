package array

import (
	"testing"
	"lib"
	"fmt"
)

func TestIntersect(t *testing.T) {
	lib.PrintFunc()

	res := Intersect([]int{1,2,2,1}, []int{2,2})

	fmt.Printf("%v\n", res)
}

func TestThreeSum(t *testing.T) {
	lib.PrintFunc()
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := ThreeSum(nums)
	fmt.Printf("%v\n", res)
}
