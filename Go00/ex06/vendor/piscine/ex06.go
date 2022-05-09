package piscine

import "ft"

func PrintCombN(n int) {
	var nums [9]int
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	for i := 0; i < n; i++ {
		ft.PrintRune('0' + rune(nums[i]))
	}
}
