package piscine

func strLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func LastRune(s string) rune {
	rs := []rune(s)
	rsLength := strLen(s)
	if rsLength == 0 {
		return 0
	}
	return rs[rsLength-1]
}
