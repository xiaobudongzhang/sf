package catalan

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	lib.PrintFunc()

	req := []int{5, 5, 5, 10, 20}
	res := LemonadeChange(req)

	fmt.Printf("%v:%v\n", req, res)

	req = []int{5, 5, 10}
	res = LemonadeChange(req)

	fmt.Printf("%v:%v\n", req, res)

	req = []int{5, 5, 10, 10, 20}
	res = LemonadeChange(req)

	fmt.Printf("%v:%v\n", req, res)
}
