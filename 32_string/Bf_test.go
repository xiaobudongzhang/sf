package _2_string

import (
	"testing"
	"lib"
	"fmt"
)

func TestBf(t *testing.T) {
	lib.PrintFunc()

	f1 := Bf("baddabef", "abc")
	f2 := Bf("baddabcef", "abc")

	fmt.Printf("f1:%v,f2:%v\n", f1, f2)
}
