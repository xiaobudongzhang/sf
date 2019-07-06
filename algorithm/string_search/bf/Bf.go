package bf

func Search(search string, pattern string) int {
	n := len(search)
	m := len(pattern)
	if m > n {
		return -1
	}
	for i :=0 ; i < (n - m +1); i++ {
		for j := 0;j < m ; j++  {
			if search[i+j] != pattern[j] {//一遍循环结束
				break
			}
			if j == m - 1 {//pattern 走完了
				return i
			}
		}
	}
	return -1
}
