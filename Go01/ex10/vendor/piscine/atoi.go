package piscine

func Atoi(s string) int {
	num := 0
	isNegative := false
	for i, c := range s {
		if i == 0 {
			if c == '+' {
				continue
			}
			if c == '-' {
				isNegative = true
				continue
			}
		}
		if c-'0' < 0 || c-'0' > 9 {
			return 0
		}
		num = num*10 + int(c-'0')
	}
	if isNegative {
		num *= -1
	}
	return num
}
