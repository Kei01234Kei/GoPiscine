package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, c := range s {
		num = num*10 + int(c-'0')
	}
	return num
}
