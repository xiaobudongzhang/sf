package _2_string

import (
	"fmt"
	"lib"
	"testing"
)

func TestGetNexts(t *testing.T) {
	lib.PrintFunc()

	nexts := GetNexts("ababacd")
	fmt.Printf("nexts:%v\n", nexts)
}

func TestKmp(t *testing.T) {
	lib.PrintFunc()

	f1 := Kmp("baddabef", "abc")
	f2 := Kmp("aaddabcef", "abc")

	fmt.Printf("f1:%v,f2:%v\n", f1, f2)

}
