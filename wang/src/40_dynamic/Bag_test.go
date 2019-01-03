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

func TestKnapsack2(t *testing.T) {
	lib.PrintFunc()

	a :=[]int{1,3,20,25,38,48,56,69,72,78}
	max := Knapsack2(a, 10, 30)
	fmt.Printf("max:%v\n", max)
}

func TestKnapsack3(t *testing.T) {
	lib.PrintFunc()

	a :=[]int{2,2,4,6,3}
	b :=[]int{3,4,8,9,6}
	max := Knapsack3(a, b,5, 9)
	fmt.Printf("max:%v\n", max)
}