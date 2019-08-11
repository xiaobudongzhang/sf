package n2

func InsertSort(a []int, n int)   {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {//不断将i插入已排序的字段
		for j := i; j > 0;j--  {

			if a[j] >= a[j-1] {
				break
			}

			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1],a[j]
			}

		}
	}
	return
}
