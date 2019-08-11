package string

import (
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	lib.PrintFunc()


	result := ""

	result = CountAndSay(1)
	if  result != "1" {
		t.Errorf("need 1, get %v",result)
	}


	result = CountAndSay(2)
	if  result != "1211" {
		t.Errorf("need 1, get %v",result)
	}


}
