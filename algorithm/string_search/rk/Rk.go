package rk

//先存储好需要的平方计算值
var powTmp = make([]int, 100)

//h[i] = 26*(h[i-1]-26^(m-1)*(s[i-1]-'a')) + (s[i+m-1]-'a');
//其中, h[i]、h[i-1] 分别对应 s[i] 和 s[i-1] 两个子串的哈希值
func Search(search string, pattern string) int {
	n := len(search)
	m := len(pattern)
	if m > n {
		return -1
	}
	makePowTmp(m)
	searchHashCode := simpleHash(search[0:m])
	patternHashCode := simpleHash(pattern)

	if patternHashCode == searchHashCode {
		return 0
	}
	for i := 1; i< (n - m +1); i++  {
		searchHashCode = (searchHashCode - (powTmp[m-1] * int(search[i-1] - 'a'))) * 26 + int(search[i+m-1] - 'a')
		if searchHashCode != patternHashCode {
			continue
		}
		if pattern != search[i:(i+m)] {
			return -1
		}
		return i
	}

	return -1
}
func makePowTmp(len int)  {
	powTmp[0] = 1
	for i:=1;i<len;i++  {
		powTmp[i] = powTmp[i-1] * 26
	}
}
//计算26进制的字符串的和
func simpleHash(subSearch string) int {
	hashCode := 0
	subLen := len(subSearch)
	for i, letter := range subSearch  {
		hashCode += int(letter - 'a') * powTmp[subLen - i-1] //subLen - i-1 倒过来算
	}
	return hashCode
}
