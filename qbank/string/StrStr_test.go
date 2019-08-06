package string

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestStrStr(t *testing.T) {
	lib.PrintFunc()

	result := -1
	haystack := ""
	needle := ""

	haystack = "hello"
	needle = "ll"

	result = StrStr(haystack, needle)
	if  result != 2 {
		t.Errorf("%v %v need 2, get %v", haystack,needle,result)
	}


	haystack = "aaaaa"
	needle = "bba"

	result = StrStr(haystack, needle)
	if  result != -1 {
		t.Errorf("%v %v need -1, get %v", haystack,needle,result)
	}

	haystack = ""
	needle = ""

	result = StrStr(haystack, needle)
	if  result != 0 {
		t.Errorf("%v %v need -1, get %v", haystack,needle,result)
	}
}
