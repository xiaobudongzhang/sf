package string

import (
	"strconv"
)

func CountAndSay(n int) string {
	myStr := "1"
	tmpStr := ""
	j := 1
	previ := -1
	itmp := 0

	for n > 0 {
		for _, i :=range myStr  {
			itmp = int(i)

			if previ == itmp {
				j++
			} else {
				j = 1
			}
			previ = itmp


			tmpStr += strconv.Itoa(j)

			tmpStr += string(itmp)
		}
		myStr = tmpStr
		tmpStr = ""
		n--
	}
	return myStr
}
