package _2_dynamic

import "math"

var minDist int
var n int
var m int
var a string
var b string

func LwstBT(a1 string, b1 string) int {
	minDist = math.MaxInt64
	n = len(a1)
	m = len(b1)
	a = a1
	b = b1
	lwstBTChild(0, 0, 0)
	return minDist
}

func lwstBTChild(i int, j int, edist int) {
	if i == n || j == m { //说明其中一个已经比较完成了

		if i < n { //a剩余的不同字符
			edist += n - i
		}

		if j < m { //b剩余的不同字符
			edist += m - j
		}

		if edist < minDist {
			minDist = edist
		}

		return
	}

	if a[i] == b[j] {
		lwstBTChild(i+1, j+1, edist)
	} else { //不匹配
		lwstBTChild(i+1, j, edist+1)   //删除a[i]或者b[j]添加一个字符
		lwstBTChild(i, j+1, edist+1)   //删除b[j]或者a[i]添加一个字符
		lwstBTChild(i+1, j+1, edist+1) //将a[i]和b[j]体会成相同字符
	}
}
