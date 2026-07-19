package problems

import (
	"fmt"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	matches := []string{}

	for a2 := range array2 {
		a2Str := array2[a2]
		for a1 := range array1 {
			a1Str := array1[a1]
			contains := strings.Contains(a2Str, a1Str)
			if contains {
				alreadyExists := false
				for _, match := range matches {
					if match == a1Str {
						alreadyExists = true
					}
				}
				if !alreadyExists {
					matches = append(matches, a1Str)
				}
			}
		}
	}

	return matches
}

func IsSolved(board [3][3]int) int {
	result := 0

	// Check straights
	for i := range board {
		rows := [3]int{board[i][0], board[i][1], board[i][2]}
		columns := [3]int{board[0][i], board[1][i], board[2][i]}
		colummnResult := checkSameNum(columns)
		rowsResult := checkSameNum(rows)
		if colummnResult != 0 {
			return colummnResult
		} else if rowsResult != 0 {
			return rowsResult
		}
	}

	diagonals := [2][3]int{
		{board[0][0], board[1][1], board[2][2]},
		{board[0][2], board[1][1], board[2][0]},
	}

	checkDiagonal1 := checkSameNum(diagonals[0])
	checkDiagonal2 := checkSameNum(diagonals[1])

	if checkDiagonal1 != 0 {
		return checkDiagonal1
	} else if checkDiagonal2 != 0 {
		return checkDiagonal2
	}

	finished := true

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				finished = false
			}
		}
	}

	if !finished {
		return -1
	}

	return result
}

func checkSameNum(arr [3]int) int {
	if arr[0] == arr[1] && arr[1] == arr[2] && arr[0] != 0 {
		if arr[0] == 1 {
			return 1
		} else {
			return 2
		}
	}
	return 0
}

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	relativeSpeed := float64(v2 - v1)
	distance := float64(g)
	time := distance / relativeSpeed
	timeInSeconds := time * 3600
	hours := int(timeInSeconds / 3600)
	minutes := int(timeInSeconds/60) % 60
	seconds := int(timeInSeconds) % 60
	return [3]int{hours, minutes, seconds}
}

func EncryptThis(text string) string {
	if len(text) == 0 {
		return ""
	}

	result := ""

	words := strings.Split(text, " ")

	for i, word := range words {
		firstChar := int(word[0])
		result += fmt.Sprintf("%d", firstChar)
		if len(word) > 1 {
			rune := []rune(word)
			rune[1], rune[len(word)-1] = rune[len(word)-1], rune[1]
			rest := string(rune)
			result += rest[1:]
		}
		if len(words) > 1 && i != len(words)-1 {
			result += " "
		}
	}

	return result
}
