package hash

import (
	//"fmt"
	"sort"
	//"bytes"
	//"strings"
	//"fmt"
	//"fmt"
	//"fmt"
	//"strings"
)

type myMap []map[string]int

var sortArr myMap

func TopKFrequent(words []string, k int) []string {
	//map统计每个url的次数
	hashmap := make(map[string]int)
	for _, word := range words {
		hashmap[word]++
	}
	//将<url,count>存储到数组中

	for k, v := range hashmap {
		sortArr = append(sortArr, map[string]int{k: v})
	}
	//排序求出需要的topk
	sort.Sort(sortArr)
	var newWords []string
	for key, item := range sortArr {
		if key+1 > k {
			break
		}
		newWords = append(newWords, getKey(item))
	}
	return newWords
}

func (this myMap) Len() int {
	return len(this)
}

func (this myMap) Less(i, j int) bool {
	if getVal(this[i]) == getVal(this[j]) { //比较字母
		return getKey(this[i])[:1] < getKey(this[j])[:1]
	}
	return getVal(this[i]) > getVal(this[j])
}

func (this myMap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func getVal(m map[string]int) int {
	for _, v := range m {
		return v
	}
	return 0
}
func getKey(m map[string]int) string {
	for k, _ := range m {
		return k
	}
	return ""
}
