package _11_sorts

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
