package _7_tree_rever

import "fmt"
//全排列，回溯思想
func Rever(arr []int, k int)  {
	if k == 1{
		for i:=0; i< len(arr);i++  {
			fmt.Printf(" %v ", arr[i])
		}
		fmt.Println()
	}
	//先排列最后一位
	for i := 0; i < k; i++ {
		arr[i], arr[k-1] = arr[k-1], arr[i]//待排列的一位的所有可能性

		Rever(arr, k -1)
		arr[i], arr[k-1] = arr[k-1], arr[i]//保证数组不变
	}

}
