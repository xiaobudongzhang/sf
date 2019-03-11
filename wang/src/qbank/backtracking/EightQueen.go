package backtracking

import "fmt"

var result [8]int//数组下标表示行，值表示列，因为一行只能有一个，所以用一维数组
var count = 0

func EightQueen(row int){
	if row == 8 {
		count++
		fmt.Printf("result: %v, count: %v\n", result, count)
		return
	}

	for i:=0; i < 8 ; i++  {
		if(canPlace(row, i)){
			result[row] = i //记录能放的位置
			//row++  //@todo 为什么不对
			EightQueen(row+1)
		}
	}
}


func canPlace(row int, col int) bool {
	left, right := col -1 , col + 1
	for i := row - 1; i>=0 ; i--  {
		//列
		if result[i] == col {
			return false
		}
		//左上
		if result[i] == left {
			return false
		}
		//右上
		if result[i] == right {
			return false
		}
		left--;
		right++;
	}
	return true
}