package bf

import (
	"fmt"
	"github.com/xiaobudongzhang/sf/lib"
	"testing"
)

func TestSearch(t *testing.T) {
	lib.PrintFunc()

	fmt.Printf("%v\n",Search("hello", "ll"))
	fmt.Printf("%v\n",Search("lldhellllod", "lll"))
	fmt.Printf("%v\n",Search("lldhellllod", "lllll"))
}
