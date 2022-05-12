package piscine

func BasicAtoi2(s string) int {
	num := 0
	for _, c := range s {
		if c-'0' < 0 || c-'0' > 9 {
			return 0
		}
		num = num*10 + int(c-'0')
	}
	return num
}
