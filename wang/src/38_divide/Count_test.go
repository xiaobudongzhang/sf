package _8_divide

import (
	"testing"
	"lib"
	"fmt"
)

func TestCount(t *testing.T) {
	lib.PrintFunc()

	num := Count([]int{2,4,3,1,5,6})
	fmt.Printf("num:%v\n", num)
}
