package greedy

import (
	"fmt"
	"lib"
	"testing"
)

func TestFindMax(t *testing.T) {
	lib.PrintFunc()

	a := []int{1, 3, 20, 25, 38, 48, 56, 69, 72, 78}
	FindMax(0, 0, a, 10, 30)
	fmt.Printf("maxW:%v\n", MaxW)
}
