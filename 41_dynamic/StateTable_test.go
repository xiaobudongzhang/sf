package _1_dynamic

import (
	"fmt"
	"lib"
	"testing"
)

func TestMinDistBt(t *testing.T) {
	lib.PrintFunc()
	MinDistBt()
}

func TestMinDistDp(t *testing.T) {
	lib.PrintFunc()
	a := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}

	min := MinDistDp(a, 4)
	fmt.Printf("min:%v\n", min)
}
