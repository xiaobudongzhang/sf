package _11_sorts

import "fmt"

func BubbleSort(a []int, n int)  {
	for i := 0; i < n; i++  {
		fmt.Printf("\nsort time:%v\n", i + 1)
		change := false
		for i := 0; i < n -1; i++  {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				change = true
			}
		}
		if !change {
			break
		}
	}

}
