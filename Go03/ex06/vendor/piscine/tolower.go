package piscine

func ToLower(s string) string {
	rs := []rune(s)
	for i, _ := range rs {
		if 'A' <= rs[i] && rs[i] <= 'Z' {
			rs[i] = rs[i] - ('A' - 'a')
		}
	}
	return string(rs)
}
