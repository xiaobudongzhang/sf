package _9_backtracking

import "fmt"

var matched bool
var plen int
var pattern string

func Match(search string, regex string) bool {
	matched = false
	plen = len(regex)
	pattern = regex

	rmatch(0, 0, search, len(search))
	return matched
}

func rmatch(ti int, pj int, text string, tlen int)  {

	if matched {//如果已经匹配了，就不需要继续递归了
		return
	}

	if pj == plen{//正则表达式到达结尾了（两个都匹配完的情况下说明匹配到了）
		if ti == tlen{//文本串到达结尾了
			matched = true
		}
		return
	}

	if pattern[pj] == '*' {//*匹配任意个字符
		for k:=0;k<=tlen-ti ;k++  {
			rmatch(ti+k, pj+1,text,tlen)
		}
	}else if pattern[pj] == '?'{//匹配0个或者一个
		fmt.Printf("??%v\n", string(pattern[pj]))
		rmatch(ti,pj+1,text,tlen)
		rmatch(ti+1,pj+1,text,tlen)
	}else if ti < tlen && pattern[pj] == text[ti] {//纯字符匹配，要完全匹配
		rmatch(ti+1, pj+1, text, tlen)
	}
}
