package _9_backtracking

import (
	"fmt"
	"lib"
	"testing"
)

func TestMatch(t *testing.T) {
	lib.PrintFunc()
	//完整字符的匹配
	res1 := Match("helloword", "he*")
	res2 := Match("helloword", "her")
	res3 := Match("helloword", "he*llo")
	res4 := Match("helloword", "he*lloword")

	fmt.Printf("res1:%v\n", res1) //tue
	fmt.Printf("res2:%v\n", res2) //false
	fmt.Printf("res3:%v\n", res3) //false
	fmt.Printf("res4:%v\n", res4) //true
}
