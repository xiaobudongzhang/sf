package _2_dynamic

import (
	"fmt"
	"lib"
	"testing"
)

func TestLwstDp(t *testing.T) {
	lib.PrintFunc()

	min := LwstDp("mitcmu", "mtacnu")
	fmt.Printf("min:%v\n", min)
}
