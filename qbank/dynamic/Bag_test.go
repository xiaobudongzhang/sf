package dynamic

import (
	"fmt"
	"lib"
	"testing"
)

func TestKnapsack(t *testing.T) {
	lib.PrintFunc()

	max := Knapsack([]int{2, 2, 4, 6, 3}, 5, 9)
	fmt.Printf("max:%v\n", max)
}
