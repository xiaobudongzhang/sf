package n2

func SelectionSort(a []int, n int)  {
	if n <= 1 {
		return
	}
	for i := 0; i < n -1; i++ {
		for j := i + 1; j < n; j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	return
}
