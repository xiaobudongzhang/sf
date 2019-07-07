package bm

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestGenerateBC(t *testing.T) {
	lib.PrintFunc()

	bc := GenerateBC("abda")
	fmt.Printf("bc:%v\n", bc)
}

func TestGenerateGS(t *testing.T) {
	lib.PrintFunc()

	suffix, prefix := GenerateGS("cabcab")
	fmt.Printf("suffix:%v,prefix:%v\n", suffix, prefix)
}

func TestSearch(t *testing.T) {
	lib.PrintFunc()

	f1 := 1//Search("baddabef", "abc")
	f2 := 1//Search("aaddabcefabcav", "abca")

	f3 := Search("System.opt.println(stack.pop())", "pop")

	fmt.Printf("f1:%v,f2:%v,f3:%v\n", f1, f2,f3)
}