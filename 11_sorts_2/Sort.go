package _11_sorts

//import "fmt"

func BubbleSort(a []int, n int) {
	for i := 0; i < n; i++ {
		change := false
		for j := 0; j < n-i-1; j++ {
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

func InsertSort(a []int, n int) {
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		tmp := a[i]
		j := i - 1
		for ; j > -1; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
			} else { //前面是有序的....
				break
			}
		}
		a[j+1] = tmp
	}
}

func SelectionSort(a []int, n int) {
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		min := a[i]
		minindex := i
		for j := i; j < n; j++ {
			if min > a[j] {
				min = a[j]
				minindex = j
			}
		}
		a[i], a[minindex] = a[minindex], a[i]
	}
}
