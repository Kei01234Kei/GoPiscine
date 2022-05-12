package piscine

func StrLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func StrRev(s string) string {
	revS := []rune(s)
	for i, j := 0, StrLen(s)-1; i < j; i, j = i+1, j-1 {
		revS[i], revS[j] = revS[j], revS[i]
	}
	return string(revS)
}
