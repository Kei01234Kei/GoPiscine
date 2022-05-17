package piscine

import "ft"

func printBoard(n int, board [8][8]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				ft.PrintRune('Q')
			} else {
				ft.PrintRune('*')
			}
		}
		ft.PrintRune('\n')
	}
	ft.PrintRune('\n')
	ft.PrintRune('\n')
}

func initBoard(board [8][8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = 0
		}
	}
}

func isQueen(n int, board [8][8]int, column int, row int, dColumn int, dRow int) bool {
	for i := 1; i < n; i++ {
		moveColumn := column + dColumn*i
		moveRow := row + dRow*i
		if moveColumn < 0 || moveColumn >= n || moveRow < 0 || moveRow >= n {
			break
		}
		if board[row+dRow*i][column+dColumn*i] == 1 {
			printBoard(n, board)
			return true
		}
	}
	return false
}

func checkQueen(n int, board [8][8]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				d := [3]int{-1, 0, 1}
				for k := 0; k < 3; k++ {
					for l := 0; l < 3; l++ {
						if d[k] == 0 && d[l] == 0 {
							continue
						}
						if isQueen(n, board, j, i, d[k], d[l]) {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

var num int

func nQueen(n int, board [8][8]int, numOfQueen int, place int) {
	if n == numOfQueen {
		if checkQueen(n, board) {
			num++
			printBoard(n, board)
		}
		return
	}
	for i := place; i < n*n; i++ {
		j := i / n
		k := i % n
		board[j][k] = 1
		if checkQueen(n, board) {
			nQueen(n, board, numOfQueen+1, place+1)
		}
		board[j][k] = 0
	}
}

func EightQueens() {
	n := 8
	var board [8][8]int
	initBoard(board)
	printBoard(n, board)
	nQueen(n, board, 0, 0)
	ft.PrintRune('0' + rune(num))
}
