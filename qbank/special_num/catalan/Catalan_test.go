package catalan

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestCataFor(t *testing.T) {
	lib.PrintFunc()

	for i := 1;i< 30;i++  {
		req := i
		res := CataFor(req)

		fmt.Printf("%v:%v\n", req, res)
	}
}
