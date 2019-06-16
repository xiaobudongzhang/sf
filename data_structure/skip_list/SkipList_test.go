package skip_list

import (
	"github.com/xiaobudongzhang/lang/golang/basic/command"
	"github.com/xiaobudongzhang/sf/lib"
	"strconv"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	command.Init()
	lib.PrintFunc()
	skipList := NewSkipList()
	for i := 1; i < 32; i++ {
		skipList.Insert("ele"+strconv.Itoa(i), float32(i))
	}
	skipList.Print()
	//skipList.Delete("ele"+strconv.Itoa(26), 26)
	//skipList.Print()
}
