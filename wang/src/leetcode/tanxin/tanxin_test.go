package tanxin

import (
	"testing"
	"lib"
	"fmt"
)

func TestFindContentChildren(t *testing.T) {
	lib.PrintFunc()

	max := FindContentChildren([]int{1,2,3}, []int{1,1})
	fmt.Printf("%v\n", max)
}