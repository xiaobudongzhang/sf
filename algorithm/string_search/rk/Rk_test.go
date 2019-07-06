package rk

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestSearch(t *testing.T) {
	lib.PrintFunc()
	if  Search("hello", "ll") != 2{
		t.Error("first err")
	}
	if Search("lldhellllod", "lll") != 5 {
		t.Error("second err")
	}
	if Search("lldhellllod", "lllll") != -1 {
		t.Error("third err")
	}
}
