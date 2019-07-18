package string

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestReverseInt(t *testing.T) {
	lib.PrintFunc()

	result := ReverseInt(123)
	if result != 321 {
		t.Error("err")
	}

	result = ReverseInt(320)
	if result != 23 {
		t.Error("err")
	}

	result = ReverseInt(-1998)
	if result != -8991 {
		t.Error("err")
	}

	result = ReverseInt(1534236469)
	if result != 0 {
		t.Error("err")
	}


}
