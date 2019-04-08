package _1_dynamic


import "fmt"

var minDist int
var minDistDp int

func MinDistBt()  {
	minDist = 99999999
	a := [][]int{
		{1,3,5,9},
		{2,1,3,4},
		{5,2,6,7},
		{6,8,4,8},
	}
	minDistBtChild(0, 0, 0,a,3 )//不包含最后一个数字
	fmt.Printf("minDist:%d\n", minDist)
}

func minDistBtChild(i int, j int, dist int,w [][]int, n int)  {
	if i == n && j == n {//走到了最后
		if dist < minDist {
			minDist = dist
		}
		return
	}

	if i < n {
		minDistBtChild(i+1, j, dist+w[i][j], w, n)
	}
	if j <n {
		minDistBtChild(i, j+1,dist+w[i][j],w,n)
	}
}

func MinDistDp(a[][]int, n int) int {
	var states [][]int
	states = make([][]int, 4)

	for i:=0;i<4 ;i++  {
		tmp := make([]int, 4)
		states[i] = tmp

		for j:=0;j<4 ;j++  {
			if i == 0 && j == 0{//初始化状态
				states[i][j] = a[i][j]
			} else if i == 0 {////初始化状态 第一行
				states[i][j] = states[i][j-1] + a[i][j]
			}else if (j==0){////初始化状态 第一列
				states[i][j] = states[i-1][j] + a[i][j]
			}else {//状态转移
				if (states[i-1][j] < states[i][j-1]){
					states[i][j] = states[i-1][j] + a[i][j]
				}else {
					states[i][j] = states[i][j-1] + a[i][j]
				}
			}
		}
	}

	return states[n-1][n-1]
}
