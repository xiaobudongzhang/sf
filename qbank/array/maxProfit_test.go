package array

import (
	"fmt"
	"lib"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	lib.PrintFunc()

	d1 := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	d2 := MaxProfit([]int{1, 2, 3, 4, 5})
	d3 := MaxProfit([]int{7, 6, 4, 3, 1})

	fmt.Printf("d1:%v\n", d1)
	fmt.Printf("d2:%v\n", d2)
	fmt.Printf("d3:%v\n", d3)
}
