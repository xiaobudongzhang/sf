package bm

import (
	"math"
)

//通过坏字符与好后缀规则，每次移动更多的位
//获取模式串中的字符位置
func GenerateBC(b string) []int {
	bc := make([]int, 256, 256)

	for k,_ := range bc {
		bc[k] = -1
	}
	for pos,letter := range b {//相同的字符后面的覆盖前面的
		bc[letter] = pos
	}
	return bc
}

func GenerateGS(pattern string) ([]int, []bool) {
	patternLen := len(pattern)
	suffix := make([]int, patternLen, patternLen)
	prefix := make([]bool, patternLen, patternLen)

	for i:=0;i<patternLen ;i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	//在[0,i]与[0,patternLen-1]中查找公共后缀，从后往前匹配,如果有相同的公共后缀，则取后面一个
	for i:=0;i < patternLen -1;i++{
		j := i
		k := 0
		for j>=0 && pattern[j] == pattern[patternLen-1-k] {//从后往前匹配
			j--
			k++
			suffix[k] = j+1//从1开始算，不是从0开始算
		}
		if j == -1 {//说明从后匹配到最开始的位置
			prefix[k] = true
		}
	}
	return suffix, prefix
}
//j 表示坏字符对应模式串的下标
func moveByGS(j int, plen int, suffix []int, prefix []bool) int {
	k := plen - j - 1 //好后缀长度
	if suffix[k] != -1 {//如果有匹配字串，则移动
		return j - suffix[k] + 1
	}
	for  r:=j+2 ;r<=plen-1;r++  {//如果有前缀匹配移动到匹配的前缀，好后缀的前缀时j+1 ，所以从j+2开始
		if prefix[plen-r] == true{
			return r
		}
	}
	return plen
}

func Search(search string, pattern string) int  {
	bc := GenerateBC(pattern)
	suffix, preffix := GenerateGS(pattern)
	n := len(search)
	m := len(pattern)

	i := 0
	for i <= n -m {
		j := -1 //坏字符在模式串中对应的下标

		for j = m -1;j>=0 ;j--  {//模式串从后往前匹配
			if search[i+j] != pattern[j] {
				break
			}
		}

		if j == -1 {//完全匹配了模式串
			return i
		}

		x := j - bc[search[i+j]] // 坏字符移动的位数,可能为负数
		y := 0
		if j < m - 1 {//如果有好后缀的话
			y = moveByGS(j, m, suffix, preffix)
		}
		i = i + int(math.Max(float64(x), float64(y)))
	}

	return -1
}
