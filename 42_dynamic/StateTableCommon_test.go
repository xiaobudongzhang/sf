package _2_dynamic

import (
	"fmt"
	"lib"
	"testing"
)

func TestLwstCommon(t *testing.T) {
	lib.PrintFunc()

	max := LwstCommon("mitcmu", "mtacnu")
	fmt.Printf("max:%v\n", max)
}
