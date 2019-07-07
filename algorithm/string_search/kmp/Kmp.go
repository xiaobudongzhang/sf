package kmp

func Search(search string, pattern string) int {
	n := len(search)
	m := len(pattern)
	nexts := GetNexts(pattern)
	j := 0
	for i := 0;i<n;i++{
		for j > 0 && search[i] != pattern[j]  {//如果不匹配,移动到最大公共前缀后一位
			j = nexts[j-1] + 1
		}
		if search[i] == pattern[j] {
			j++
		}
		if j == m {//走完了模式串
			return i - m + 1 //i从0开始, 所以加1
		}
	}
	return -1
}
/**
 *
* 最难理解的地方是
* k = next[k]
* 因为前一个的最长串的下一个字符不与最后一个相等，
* 需要找前一个的次长串，
* 问题就变成了求0到next(k)的最长串，
* 如果下个字符与最后一个不等，继续求次长串，
* 也就是下一个next(k)，直到找到，或者完全没有
 */
func GetNexts(pattern string) []int  {
	patternLen := len(pattern)
	nexts := make([]int, patternLen, patternLen)
	nexts[0] = -1
	k := -1

	for i := 1; i < patternLen;i++  {
		for k != -1 && pattern[k+1] != pattern[i]  {
			k = nexts[k]
		}
		if pattern[k+1] == pattern[i] {
			k++
		}
		nexts[i] = k
	}
	return nexts
}