package _9_backtracking

import "math"

func MaxProfit(prices []int) int {
	min_p,max_p := math.MaxInt32, 0
	for i:=0;i<len(prices);i++{
		min_p = int(math.Min(float64(min_p), float64(prices[i])))
		max_p = int(math.Max(float64(max_p), float64(prices[i] - min_p)))
	}
	return max_p
}
