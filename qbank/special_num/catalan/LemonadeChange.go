package catalan

func LemonadeChange(bills []int) bool {

	var five int
	var ten int

	for _, bill := range bills {
		if bill == 5 {
			five++
			continue
		}
		if bill == 10 {
			ten++
			five--
			if five < 0 {
				return false
			}
			continue
		}
		//20 = 10 + 5 or 5+5+5
		if ten > 0 && five > 0 {
			ten--
			five--
			continue
		}
		if five >= 3 {
			five = five - 3
			continue
		}
		return false
	}
	return true
}
