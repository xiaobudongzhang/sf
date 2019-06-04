package _2_dynamic

func LwstCommon(a string, b string) int {
	var states [][]int
	states = make([][]int, len(a))

	for i := 0; i < len(a); i++ {
		tmp := make([]int, len(b))
		states[i] = tmp

		for j := 0; j < len(b); j++ {
			if i == 0 && j == 0 { //初始化状态

				if a[i] == b[j] {
					states[i][j] = 1
				} else {
					states[i][j] = 0
				}

			} else if i == 0 { ////初始化状态 第一行

				if a[i] == b[j] {
					states[i][j] = states[i][j-1] + 1
				} else {
					states[i][j] = states[i][j-1]
				}

			} else if j == 0 { ////初始化状态 第一列

				if a[i] == b[j] {
					states[i][j] = states[i-1][j] + 1
				} else {
					states[i][j] = states[i-1][j]
				}

			} else { //状态转移

				if a[i] == b[j] {
					states[i][j] = max(states[i-1][j], states[i][j-1], states[i-1][j-1]+1)
				} else {
					states[i][j] = max(states[i-1][j], states[i][j-1], states[i-1][j-1])
				}

			}
		}
	}
	//找到最大值
	max := -1
	aindex := len(a) - 1
	for j := 0; j < len(b); j++ {
		if states[aindex][j] > max {
			max = states[aindex][j]
		}
	}
	return max
}

func max(a, b, c int) int {
	max := -1
	if a > max {
		max = a
	}

	if b > max {
		max = b
	}

	if c > max {
		max = c
	}
	return max
}
