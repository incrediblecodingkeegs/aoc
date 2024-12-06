package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")

	fmt.Printf("Part One: %d\n", partOne(string(f)))
	fmt.Printf("Part Two: %d\n", partTwo(string(f)))

}

func partOne(s string) int {
	rows := strings.Split(s, "\n")
	cols := getCols(rows)
	mainDiags, antiDiags := getDiags(rows, cols)
	//fmt.Printf("rows: %s\ncols: %s\nmainDiags: %s\nantiDiags: %s\n", rows, cols, mainDiags, antiDiags)
	return countXmas(rows) + countXmas(cols) + countXmas(mainDiags) + countXmas(antiDiags)
}

func partTwo(s string) int {
	rows := strings.Split(s, "\n")

	result := 0
	for i := 1; i < len(rows)-1; i++ {
		row := rows[i]
		for j := 1; j < len(row)-1; j++ {
			if row[j] == 'A' {
				if isXmas(rows, i, j) {
					result++
				}
			}
		}
	}

	return result

}

func isXmas(rows []string, x, y int) bool {
	// check if corners around A have two 'S's and 'M's.
	var corners string
	corners += string(rows[x-1][y+1]) // top left corner
	corners += string(rows[x-1][y-1]) // bottom left corner
	corners += string(rows[x+1][y-1]) // bottom right corner
	corners += string(rows[x+1][y+1]) // top right corner

	// needs exactly two 'S's and 'M's, and they need to be on touching corners to make an X shape
	fmt.Println("corners ", corners)
	if strings.Contains(corners, "SS") || strings.Contains(corners, "MM") {
		if strings.Count(corners, "S") == 2 && strings.Count(corners, "M") == 2 {
			return true
		}
	}
	return false
}

func getCols(rows []string) []string {
	cols := make([]string, len(rows[0]))
	for _, row := range rows {
		for j := range row {
			cols[j] += string(row[j])
		}
	}
	return cols
}

func getDiags(rows, cols []string) ([]string, []string) {

	var mainDiagonals []string
	var antiDiagonals []string

	// left to right diagonals

	// \ \ \
	// . \ \
	// . . \
	for col := 0; col < len(cols); col++ {
		var diagonal string
		for i, j := 0, col; i < len(rows) && j < len(cols); i, j = i+1, j+1 {
			diagonal += string(rows[i][j])
		}
		mainDiagonals = append(mainDiagonals, diagonal)
	}

	// . . .
	// \ . .
	// \ \ .
	for row := 1; row < len(rows); row++ {
		var diagonal string
		for i, j := row, 0; i < len(rows) && j < len(cols); i, j = i+1, j+1 {
			diagonal += string(rows[i][j])
		}
		mainDiagonals = append(mainDiagonals, diagonal)
	}

	//fmt.Println(mainDiagonals)

	// right to left diagonals

	// / / /
	// / / .
	// / . .
	for col := len(cols) - 1; col >= 0; col-- {
		var diagonal string
		for i, j := 0, col; i < len(rows) && j >= 0; i, j = i+1, j-1 {
			diagonal += string(rows[i][j])
		}
		antiDiagonals = append(antiDiagonals, diagonal)
	}

	// . . .
	// . . /
	// . / /
	for row := 1; row <= len(rows); row++ {
		var diagonal string
		for i, j := row, len(cols)-1; i < len(rows) && j >= 0; i, j = i+1, j-1 {
			diagonal += string(rows[i][j])
		}
		antiDiagonals = append(antiDiagonals, diagonal)
	}
	//fmt.Println(antiDiagonals)

	return mainDiagonals, antiDiagonals
}

func countXmas(strs []string) int {
	count := 0
	for _, str := range strs {
		count += strings.Count(str, "XMAS") + strings.Count(str, "SAMX")
	}
	return count
}
