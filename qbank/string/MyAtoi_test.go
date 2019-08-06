package string

import (
	"github.com/xiaobudongzhang/sf/lib"
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	lib.PrintFunc()
	result := MyAtoi("42")
	if result != 42 {
		t.Error("need 42 ,get" , result)
	}

	result = MyAtoi("   -42")
	if result != -42 {
		t.Error("need -42 , get ", result)
	}

	result = MyAtoi("4193 with words")
	if result != 4193 {
		t.Error("need 4193 , get ", result)
	}

	result = MyAtoi("words and 987")
	if result != 0 {
		t.Error("need 0 ,get", result)
	}

	result = MyAtoi( "-91283472332")
	if result != -1 * int(math.Pow(2, 31)) {
		t.Error("need min, get", result)
	}

	result = MyAtoi( "+1")
	if result != 1 {
		t.Error("need 1, get", result)
	}
	result = MyAtoi( "+-2")
	if result != 0 {
		t.Error("need 0, get", result)
	}

	result = MyAtoi( "   +0 123")
	if result != 0 {
		t.Error("need 0, get", result)
	}

	result = MyAtoi( "9223372036854775808")
	if result != 2147483647 {
		t.Error("need 2147483647, get", result)
	}

	result = MyAtoi( "-91283472332")
	if result != -2147483648 {
		t.Error("need 2147483647, get", result)
	}
	result = MyAtoi( "0-1")
	if result != 0 {
		t.Error("need 0, get", result)
	}

	result = MyAtoi( "0  123")
	if result != 0 {
		t.Error("need 0, get", result)
	}
	result = MyAtoi( "-   234")
	if result != 0 {
		t.Error("need 0, get", result)
	}
	result = MyAtoi( "    -88827   5655  U")
	if result != -88827 {
		t.Error("need 0, get", result)
	}
	result = MyAtoi( "-5-")
	if result != -5 {
		t.Error("need 0, get", result)
	}

}
