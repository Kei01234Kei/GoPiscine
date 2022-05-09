package piscine

import "ft"

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= j {
				continue
			}
			for k := 0; k < 10; k++ {
				if j >= k {
					continue
				}
				err := ft.PrintRune('0' + rune(i))
				if err != nil {
					return
				}
				err = ft.PrintRune('0' + rune(j))
				if err != nil {
					return
				}
				err = ft.PrintRune('0' + rune(k))
				if err != nil {
					return
				}
				if i == 7 && j == 8 && k == 9 {
					return
				}
				err = ft.PrintRune(',')
				if err != nil {
					return
				}
			}
		}
	}
}
