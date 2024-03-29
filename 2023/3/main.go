package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")

	total := 0
	for _, partNumber := range findPartNumbers(lines) {
		total += partNumber
	}
	fmt.Printf("total = %d\n", total)
}

func findPartNumbers(schematic []string) []int {
	re, _ := regexp.Compile(`(\d+)`)
	var partNumbers []int
	for i, line := range schematic {
		numIndexes := re.FindAllStringIndex(line, -1)

		// build box around string
		leftIndex, rightIndex := 0, 0
		for _, numIndex := range numIndexes {
			var substringSlice []string
			// get left index
			if numIndex[0] != 0 {
				leftIndex = numIndex[0] - 1
			} else {
				leftIndex = numIndex[0]
			}
			// get right index
			if numIndex[1] == len(line) {
				rightIndex = numIndex[1]
			} else {
				rightIndex = numIndex[1] + 1
			}
			//fmt.Printf("len(line) = %d leftIndex = %d rightIndex = %d\n", len(line), leftIndex, rightIndex)
			// add substring in line before match if it exists
			if i != 0 {
				substringSlice = append(substringSlice, schematic[i-1][leftIndex:rightIndex])
			}
			substringSlice = append(substringSlice, schematic[i][leftIndex:rightIndex])
			// add substring in line after match if it exists
			if i != len(schematic)-1 {
				substringSlice = append(substringSlice, schematic[i+1][leftIndex:rightIndex])
			}
			if isPartNumber(substringSlice) {
				partNumber, err := strconv.Atoi(line[numIndex[0]:numIndex[1]])
				fmt.Printf("partNumber = %d\n", partNumber)
				if err != nil {
					fmt.Println(err)
				}
				partNumbers = append(partNumbers, partNumber)
			}
		}

	}
	return partNumbers
}

func isPartNumber(substringSlice []string) bool {
	fmt.Printf("%v", substringSlice)
	for _, substring := range substringSlice {
		for _, r := range substring {
			if !unicode.IsDigit(r) && r != '.' {
				fmt.Println(true)
				return true
			}
		}
	}
	fmt.Println(false)

	return false
}
