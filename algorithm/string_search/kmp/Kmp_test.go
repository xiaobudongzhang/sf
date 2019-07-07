package kmp

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestGetNexts(t *testing.T) {
	lib.PrintFunc()

	nexts := GetNexts("ababacd")
	fmt.Printf("nexts:%v\n", nexts)
}

func TestSearch(t *testing.T) {
	lib.PrintFunc()

	f1 := Search("baddabef", "abc")
	f2 := Search("aaddabcef", "abc")

	fmt.Printf("f1:%v,f2:%v\n", f1, f2)

}
