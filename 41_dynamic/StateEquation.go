package _1_dynamic

import (
	//"math"
	//"fmt"
	"lib"
)

//min_dist(i,j)= w[i][j] + min(min_dist(i-1,j),min_dist(i,j-1))
var datas [][]int

var states [][]int

func MindistP()  int{
	datas =[][]int{
		{1,3,5,9},
		{2,1,3,4},
		{5,2,6,7},
		{6,8,4,3},
	}
	states =[][]int{
		{-1,-1,-1,-1},
		{-1,-1,-1,-1},
		{-1,-1,-1,-1},
		{-1,-1,-1,-1},
	}
	return Mindist(3,3)
}

func Mindist(i int, j int) int {
	if i==0 && j == 0 {
		return datas[i][j]
	}
	if states[i][j] > 0 {
		return states[i][j]
	}
	left := 88888888888888
	if i > 0 {
		left = Mindist(i-1, j)
	}
	right := 88888888888888
	if j >0 {
		right = Mindist(i, j-1)
	}

	cur := lib.MinInt(left, right) + datas[i][j]
	states[i][j] = cur
	return cur
}
