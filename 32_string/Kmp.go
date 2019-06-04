package _2_string

func Kmp(a string, p string) int {

	alen := len(a)
	plen := len(p)
	next := GetNexts(p)

	j := 0
	for i := 0; i < alen-plen+1; i++ {

		for j > 0 && a[i] != p[j] { //>0 有前缀(kmp的关键点)
			j = next[j-1] + 1 //next[j-1]是最长匹配前缀 所以加1走到不匹配的第一个位置
		}

		if a[i] == p[j] {
			j++
		}

		if j == plen {
			return i - plen + 1
		}

	}

	return -1
}

func GetNexts(p string) []int {
	plen := len(p)
	next := make([]int, plen)

	next[0] = -1
	k := -1

	for i := 1; i < plen; i++ {

		for k != -1 && p[k+1] != p[i] {
			k = next[k]
		}

		if p[k+1] == p[i] { //如果最大匹配后缀字符的下一个字符跟字符串的下一个字符相同，则下一个匹配串的最大可匹配子串是最大可匹配子串加一
			k++
		}

		next[i] = k
	}

	return next
}
