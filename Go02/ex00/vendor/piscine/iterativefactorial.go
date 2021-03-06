package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	intMax := int(^uint(0) >> 1)
	num := 1
	for nb > 0 {
		if intMax/nb < num {
			return 0
		}
		num *= nb
		nb--
	}
	return num
}
