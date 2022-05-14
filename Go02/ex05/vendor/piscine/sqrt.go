package piscine

func Sqrt(nb int) int {
	i := 1
	for i <= nb/i {
		if i == nb/i {
			return i
		}
		i++
	}
	return 0
}
