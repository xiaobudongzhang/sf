package _8_stack

import (
	"fmt"
	"testing"
)

func TestBrowser_Prev(t *testing.T) {
	PrintFunc()
	browser := NewBrowser()
	browser.PushPrev("a")
	browser.PushPrev("b")
	browser.PushPrev("c")

	fmt.Printf("prev:%v,next:%v\n", browser.prev, browser.next)

	browser.Next()
	browser.Next()
	fmt.Printf("prev:%v,next:%v\n", browser.prev, browser.next)

	browser.Prev()
	fmt.Printf("prev:%v,next:%v\n", browser.prev, browser.next)

	browser.Open("d")
	fmt.Printf("prev:%v,next:%v\n", browser.prev, browser.next)

	browser.Open("e")
	fmt.Printf("prev:%v,next:%v\n", browser.prev, browser.next)
}
