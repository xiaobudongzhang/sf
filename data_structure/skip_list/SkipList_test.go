package skip_list

import (
	"github.com/xiaobudongzhang/sf/lib"
	"strconv"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	lib.PrintFunc()
	skipList := NewSkipList()
	for i:=1;i<50 ;i++  {
		skipList.Insert("ele" + strconv.Itoa(i), float32(i))
	}
	skipList.Print()
	skipList.Delete("ele" + strconv.Itoa(26), 26)
	skipList.Print()
}


