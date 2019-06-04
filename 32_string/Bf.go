package _2_string

func Bf(search string, pattern string) int {

	patternlen := len(pattern)

	for i := 0; i < len(search)-patternlen; i++ {
		if search[i:(i+patternlen)] == pattern {
			return i
		}
	}
	return -1
}
