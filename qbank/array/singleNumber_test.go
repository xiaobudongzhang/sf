package array

import (
	"fmt"
	"lib"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	lib.PrintFunc()

	res := SingleNumber([]int{4, 1, 2, 1, 2})
	fmt.Printf("%v\n", res)
}
