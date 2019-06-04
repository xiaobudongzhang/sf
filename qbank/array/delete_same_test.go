package array

import (
	"fmt"
	"lib"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	lib.PrintFunc()

	xx := []int{1, 1, 2}
	len := RemoveDuplicates(xx)
	fmt.Printf("%v:%v", len, xx[:len])
}
