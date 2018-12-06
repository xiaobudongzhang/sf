package _8_stack

import (
	"testing"
	"fmt"
	"runtime"
)

func TestArrayStack_Push(t *testing.T) {
	PrintFunc()
	stack := NewArrayStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	fmt.Printf("%v:%v\n", stack.stack, stack.top)
}

func TestArrayStack_Pop(t *testing.T) {
	PrintFunc()
	stack := NewArrayStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)


	fmt.Printf("%v:%v\n", stack.Pop(), stack.top)
}

func TestArrayStack_Top(t *testing.T) {
	PrintFunc()
	stack := NewArrayStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Printf("%v:%v\n", stack.stack, stack.Top())
}

func PrintFunc()  {
	fmt.Println("--------------------------------------")
	funName, _,_,ok := runtime.Caller(1)
	if ok {
		fmt.Println("funName+" + runtime.FuncForPC(funName).Name())
	}
}