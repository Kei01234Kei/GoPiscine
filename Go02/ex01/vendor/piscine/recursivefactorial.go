package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	num := 1
	num = nb * RecursiveFactorial(nb-1)
	return num
}