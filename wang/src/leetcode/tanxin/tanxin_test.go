package tanxin

import (
	"testing"
	"lib"
	"fmt"
)

func TestFindContentChildren(t *testing.T) {
	lib.PrintFunc()

	max := FindContentChildren([]int{7,8}, []int{3,4,5,6,7,8})
	fmt.Printf("%v\n", max)
}