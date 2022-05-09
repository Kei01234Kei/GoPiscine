package main

import "ft"

func main() {
	char := 'z'
	for i := 0; i < 26; i++ {
		err := ft.PrintRune(char - rune(i))
		if err != nil {
			return
		}
	}
	err := ft.PrintRune('\n')
	if err != nil {
		return
	}
}
