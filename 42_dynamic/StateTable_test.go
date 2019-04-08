package _2_dynamic

import (
	"testing"
	"lib"
	"fmt"
)

func TestLwstDp(t *testing.T) {
	lib.PrintFunc()

	min := LwstDp("mitcmu", "mtacnu")
	fmt.Printf("min:%v\n", min)
}
