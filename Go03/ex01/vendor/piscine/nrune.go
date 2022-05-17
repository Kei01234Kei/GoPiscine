package piscine

func strLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func NRune(s string, n int) rune {
	rs := []rune(s)
	rsLength := strLen(s)
	if n <= 0 || n > rsLength {
		return 0
	}
	return rs[n-1]
}
