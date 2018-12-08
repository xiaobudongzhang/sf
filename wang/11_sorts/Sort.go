package _11_sorts

import "fmt"

func BubbleSort(a []int, n int)  {
	for i := 0; i < n; i++  {
		change := false
		for j := 0; j < n - i -1; j++  {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
				change = true
			}
		}
		if !change {
			break
		}
	}

}

func InsertSort(a []int, n int)  {
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		tmp := a[i]
		for j := i -1; j > -1; j-- {
			fmt.Printf("j:%v", j)
			if tmp < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
