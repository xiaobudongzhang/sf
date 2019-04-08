package _2_dynamic

import (
	"testing"
	"lib"
	"fmt"
)

func TestLwstBT(t *testing.T) {
	lib.PrintFunc()

	dif := LwstBT("mitcmu", "mtacnu")
	fmt.Printf("dif:%v\n", dif)
}