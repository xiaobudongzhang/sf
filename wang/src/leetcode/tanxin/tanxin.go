package tanxin

import (
	"sort"
	"fmt"
)

func FindContentChildren(g []int, s []int) int {
	//找出对糖果需求最小的，然后找出能满足的最小的饼干
	sort.Ints(g)
	sort.Ints(s)
	fmt.Printf("%v,%v\n", g, s)
	max, i, j := 0, 0, 0
	for i< len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			max++
		}
		j++
	}
	return max
}