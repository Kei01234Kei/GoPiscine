package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	num := 1
	num = num * nb * RecursivePower(nb, power-1)
	return num
}
