package _2_string

import (
	"fmt"
	"math"
)

func GenerateBC(b string) []int {
	bc := make([]int, 256)

	for i := 0; i < 256; i++ {
		bc[i] = -1
	}
	for i, c := range b {
		bc[c] = i
	}
	return bc
}

//坏字符
func BadChar(a string, pattern string) int {
	bc := GenerateBC(pattern)
	i := 0
	alen := len(a)

	patternlen := len(pattern)
	for i < alen-patternlen {
		j := -1
		for j = patternlen - 1; j >= 0; j-- {
			if pattern[j] != a[i+j] { //坏字符
				break
			}
		}
		if j < 0 {
			return i
		}

		i = i + (j - bc[a[i+j]])
	}
	return -1
}

func GenerateGS(pattern string) ([]int, []bool) {
	plen := len(pattern)
	suffix := make([]int, plen)
	prefix := make([]bool, plen)
	for i := 0; i < plen; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < plen-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[plen-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
	return suffix, prefix
}

func Bm(a string, p string) int { //所有的规则都是为了移动更多位置
	bc := GenerateBC(p)             //坏字符hash表，计算坏字符最后出现的位置
	suffix, prefix := GenerateGS(p) //好后缀字子字符表

	alen := len(a)
	plen := len(p)
	i := 0

	for i <= alen-plen {
		//坏字符规则
		j := -1
		for j = plen - 1; j >= 0; j-- {
			if p[j] != a[i+j] { //坏字符
				break
			}
		}
		if j < 0 {
			return i
		}
		//根据j判断是否有好后缀
		x := j - bc[a[i+j]] //坏字符移动的位数
		y := 0
		if j < plen-1 { //有好后缀
			y = moveByGS(j, plen, suffix, prefix)
		}
		i = i + int(math.Max(float64(x), float64(y)))
	}

	return -1
}

func moveByGS(j int, plen int, suffix []int, prefix []bool) int {
	k := plen - 1 - j
	//1.好后缀有匹配
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	//好后缀的子串有匹配的前缀
	for r := j + 2; r < plen; r++ {
		if prefix[plen-r] == true {
			fmt.Printf("x:%v", r)
			return r
		}
	}
	//完全不匹配，说明可以直接移动整个模式串长度
	return plen
}
