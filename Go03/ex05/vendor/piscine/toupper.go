package piscine

func ToUpper(s string) string {
	rs := []rune(s)
	for i, _ := range rs {
		if 'a' <= rs[i] && rs[i] <= 'z' {
			rs[i] = 'A' - 'a' + rs[i]
		}
	}
	return string(rs)
}
