package dynamic

import (
	"fmt"
	"lib"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	lib.PrintFunc()

	fmt.Printf("%v\n", MaxSubArray([]int{2, 1, -5, 4})) //-2,1,-3,4,-1,2,1,-5,4
}

func TestRob(t *testing.T) {
	lib.PrintFunc()

	fmt.Printf("%v\n", Rob([]int{2, 1, 1, 2})) //-2,1,-3,4,-1,2,1,-5,4
}
