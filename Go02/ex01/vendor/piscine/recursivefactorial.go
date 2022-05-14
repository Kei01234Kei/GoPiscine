package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if 0 <= nb && nb <= 1 {
		return 1
	}
	intMax := int(^uint(0) >> 1)
	num := 1
	if intMax/(nb-1) < nb {
		return 0
	}
	num = nb * RecursiveFactorial(nb-1)
	if intMax/nb < num {
		return 0
	}
	return num
}
