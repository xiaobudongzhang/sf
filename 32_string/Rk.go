package _2_string

import (
	//"strconv"
	"math"
	//"fmt"
	//"fmt"
	//"unicode/utf8"
)

//h[i] = 26*(h[i-1]-26^(m-1)*(s[i-1]-'a')) + (s[i+m-1]-'a');
//其中, h[i]、h[i-1] 分别对应 s[i] 和 s[i-1] 两个子串的哈希值

func Rk(search string, pattern string) int {

	patternlen := len(pattern)
	patternhash := simpleHash(pattern)

	var decimal int64
	decimal = 128 //进制

	h := make([]int64, len(search) - patternlen)
	h[0] = simpleHash(search[0:patternlen])
	//求hash
	for i:=1; i< len(search) - patternlen;i++  {
		pow := int64(math.Pow(float64(decimal), float64(patternlen-1)))
		h[i] = decimal * (h[i-1]-  pow * simpleHash(search[i-1:i])) + simpleHash(search[i+patternlen-1:i+patternlen])
	}

	//根据hash比较
	for i:=0;i<len(search) - patternlen;i++  {
		if h[i] != patternhash {
			continue
		}
		if search[i:i+patternlen] == pattern {
			return i
		}
	}
	return -1
}

func simpleHash(str string) int64 {
	var decimal int64
	decimal = 128 //进制
	var hash int64
	hash = 0
	len := len(str)
	for i, c := range str {//获取字符的asscii数字用range  不要用for 这是一个坑
		powlen := int64(len) - int64(i + 1)
		pow := int64(math.Pow(float64(decimal), float64(powlen)))
		hash += int64(c) * pow
	}
	return (hash)
}
