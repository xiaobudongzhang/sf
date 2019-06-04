package _2_dynamic

import (
	"fmt"
	"lib"
	"testing"
)

func TestLwstBT(t *testing.T) {
	lib.PrintFunc()

	dif := LwstBT("mitcmu", "mtacnu")
	fmt.Printf("dif:%v\n", dif)
}
