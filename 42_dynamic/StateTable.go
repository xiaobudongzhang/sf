package _2_dynamic

import "math"

var states [][]int

func LwstDp(a string, b string) int {
	var states [][]int
	states = make([][]int, len(a))

	for i := 0; i < len(a); i++ {
		tmp := make([]int, len(b))
		states[i] = tmp

		for j := 0; j < len(b); j++ {
			if i == 0 && j == 0 { //初始化状态

				if a[i] == b[j] {
					states[i][j] = 0
				} else {
					states[i][j] = 1
				}

			} else if i == 0 { ////初始化状态 第一行

				if a[i] == b[j] {
					states[i][j] = states[i][j-1]
				} else {
					states[i][j] = states[i][j-1] + 1
				}

			} else if j == 0 { ////初始化状态 第一列

				if a[i] == b[j] {
					states[i][j] = states[i-1][j]
				} else {
					states[i][j] = states[i-1][j] + 1
				}

			} else { //状态转移

				if a[i] == b[j] {
					states[i][j] = min(states[i-1][j]+1, states[i][j-1]+1, states[i-1][j-1])
				} else {
					states[i][j] = min(states[i-1][j]+1, states[i][j-1]+1, states[i-1][j-1]+1)
				}

			}
		}
	}

	//找到最小值
	min := math.MaxInt32
	aindex := len(a) - 1
	for j := 0; j < len(b); j++ {
		if states[aindex][j] < min {
			min = states[aindex][j]
		}
	}
	return min
}

func min(a, b, c int) int {
	min := math.MaxInt32
	if a < min {
		min = a
	}

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}
	return min
}
