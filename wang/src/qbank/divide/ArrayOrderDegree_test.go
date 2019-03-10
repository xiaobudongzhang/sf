package divide

import (
	"testing"
	"lib"
	"fmt"
)

func TestArrayOrderDegree(t *testing.T) {
	lib.PrintFunc()
	arr := []int{1,5,6,2,4,3}
	count := ArrayOrderDegree(arr)
	fmt.Printf("%v\n", count)
	fmt.Printf("%v\n", arr)
}
