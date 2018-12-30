package _2_string

import (
	"testing"
	"lib"
	"fmt"
)

func TestGenerateBC(t *testing.T) {
	lib.PrintFunc()

	bc := GenerateBC("abda")
	fmt.Printf("bc:%v\n", bc)
}

func TestBadChar(t *testing.T) {
	lib.PrintFunc()

	index := BadChar("abcacabdc", "abd")
	fmt.Printf("abd index:%v\n", index)
}

func TestGenerateGS(t *testing.T) {
	lib.PrintFunc()

	suffix, prefix := GenerateGS("cabcab")
	fmt.Printf("suffix:%v,prefix:%v\n", suffix,prefix)
}

func TestBm(t *testing.T) {
	lib.PrintFunc()

	f1 := Bf("baddabef", "abc")
	f2 := Bf("baddabcef", "abc")

	fmt.Printf("f1:%v,f2:%v\n", f1, f2)
}