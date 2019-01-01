package _9_backtracking

import "fmt"
var result [8]int
var count = 0
//result = make([]int,8,8)//数组下标代表行 数组值代表列

func Cal8Queue(row int)  {
	if row == 8 {
		count++
		fmt.Printf("count:%v\n", count)
		printQueens(result)
		return
	}
	for i:=0;i<8 ;i++  {//每一列进行判断
		if isOk(row, i){
			result[row] = i
			//row++ //错误用法，相当于多加了一次所以直到7
			Cal8Queue(row+1)
		}
	}
}

func isOk(row int, column int) bool {

	left,right := column -1, column + 1
	for i:=row-1;i>=0 ;i--  {
		//1.当前列
		if result[i] == column {
			return  false
		}
		//2.对角线位置(两个必须都符合才行)
		if left >=0 && result[i] == left {
			return false
		}
		if right < 8 && result[i] == right {
			return false
		}
		left--
		right++
	}
	return true
}

func printQueens(result [8]int)  {
	for i:=0;i<8 ;i++  {//行
		for j:=0;j<8 ;j++  {//列
			if result[i] == j {
				fmt.Printf("q")
			}else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
