package piscine

import "ft"

func IsNegative(nb int) {
	if nb < 0 {
		err := ft.PrintRune('T')
		if err != nil {
			return
		}
	} else {
		err := ft.PrintRune('F')
		if err != nil {
			return
		}
	}
	err := ft.PrintRune('\n')
	if err != nil {
		return
	}
}
