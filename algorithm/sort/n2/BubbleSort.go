package n2

func BubbleSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 0;i<n ;i++  {
		flag := false
		for j := 0;j<n -1;j++  {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
	return a
}
