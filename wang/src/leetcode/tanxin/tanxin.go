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
	max := 0
	for _, gx := range g {

		if len(s) < 1{
			break
		}
		//找到大于等于所需需求的饼干
		for _,sx := range s {
			if len(s) == 1 {
				s = nil
			}else {
				s = s[1:]
			}

			if sx >= gx {
				max++
				break
			}
		}
	}

	return max
}