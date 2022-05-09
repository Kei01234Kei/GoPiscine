package piscine

import "ft"

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i >= j {
				continue
			}
			onesPlaceOfI := i % 10
			tensPlaceOfI := i / 10
			if i == 0 {
				err := ft.PrintRune('0')
				if err != nil {
					return
				}
				err = ft.PrintRune('0')
				if err != nil {
					return
				}
				err = ft.PrintRune(' ')
				if err != nil {
					return
				}
			} else {
				err := ft.PrintRune('0' + rune(tensPlaceOfI))
				if err != nil {
					return
				}
				err = ft.PrintRune('0' + rune(onesPlaceOfI))
				if err != nil {
					return
				}
				err = ft.PrintRune(' ')
				if err != nil {
					return
				}
			}
			onesPlaceOfJ := j % 10
			tensPlaceOfJ := j / 10
			err := ft.PrintRune('0' + rune(tensPlaceOfJ))
			if err != nil {
				return
			}
			err = ft.PrintRune('0' + rune(onesPlaceOfJ))
			if err != nil {
				return
			}
			if i != 98 || j != 99 {
				err = ft.PrintRune(',')
				if err != nil {
					return
				}
				err = ft.PrintRune(' ')
				if err != nil {
					return
				}
			} else {
				err := ft.PrintRune('\n')
				if err != nil {
					return
				}
			}
		}
	}
}
