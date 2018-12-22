package _7_tree_rever

import "fmt"

func Rever(arr []int, k int)  {
	if k == 1{
		for i:=0; i< len(arr);i++  {
			fmt.Printf(" %v ", arr[i])
		}
		fmt.Println()
	}

	for i := 0; i < k; i++ {
		arr[i], arr[k-1] = arr[k-1], arr[i]

		Rever(arr, k -1)
		arr[i], arr[k-1] = arr[k-1], arr[i]
	}

}
