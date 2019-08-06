package string

func StrStr(haystack string, needle string) int {
	if needle == "" && haystack == "" {
		return 0
	}
	if needle == "" {
		return 0
	}

	hlen := len(haystack)
	nlen := len(needle)

	for i := 0;i< hlen-nlen +1 ; i++  {
		for j:=0;j<nlen ;j++  {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == nlen -1 {
				return i
			}
		}
	}
	return -1
}
