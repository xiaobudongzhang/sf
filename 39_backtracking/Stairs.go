package _9_backtracking

import "fmt"

var count int

func ClimbStairs(n int) int {
	count = 0
	climbStairsChild(n, 0)
	fmt.Printf("%v\n", count)
	return count
}

func climbStairsChild(n int, step int) {
	fmt.Printf("%v-%v\n", step, n)
	if step > n {
		return
	}
	if step == n {
		count++
		return
	}
	//有两种走法
	if step+1 <= n {
		climbStairsChild(n, step+1)
	}
	if step+2 <= n {
		climbStairsChild(n, step+2)
	}
}
