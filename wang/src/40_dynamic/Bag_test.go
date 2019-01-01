package _0_dynamic

import (
	"testing"
	"lib"
	"fmt"
)

func TestKnapsack(t *testing.T) {
	lib.PrintFunc()

	a :=[]int{1,3,20,25,38,48,56,69,72,78}
	max := Knapsack(a, 10, 24)
	fmt.Printf("max:%v\n", max)
}
