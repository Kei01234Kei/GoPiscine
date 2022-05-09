package main

import (
	"ft"
)

func main() {
	for i := 48; i < 58; i++ {
		err := ft.PrintRune(rune(i))
		if err != nil {
			return
		}
	}
	err := ft.PrintRune('\n')
	if err != nil {
		return
	}
}
