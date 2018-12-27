package _2_string

import (
	"testing"
	"lib"
	"fmt"
)

func TestRk(t *testing.T) {
	lib.PrintFunc()

	f1 := Rk("baddabef", "abc")
	f2 := Rk("baddabcef", "abc")

	fmt.Printf("f1:%v,f2:%v\n", f1, f2)
}
