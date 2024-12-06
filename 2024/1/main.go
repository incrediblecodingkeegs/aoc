package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")

	var col1, col2 []int

	for i, line := range lines {
		// split on all whitespace
		ints := strings.Fields(line)
		if len(ints) != 2 {
			log.Fatalf("invalid input.txt, expected 2 got %d, line: %d - %s", len(ints), i, line)
		}
		int0, _ := strconv.Atoi(ints[0])
		int1, _ := strconv.Atoi(ints[1])

		// parse each field into an integer,
		col1 = append(col1, int0)
		col2 = append(col2, int1)
	}

	result := partOne(col1, col2)
	fmt.Println("Part1 Result: ", result)

	result = partTwo(col1, col2)
	fmt.Println("Part2 Result: ", result)

}

func partOne(col1, col2 []int) int {
	// sort each array
	sort.Ints(col1)
	sort.Ints(col2)

	// sum the differences between them
	result := 0
	for i := 0; i < len(col1); i++ {
		result += int(math.Abs(float64(col1[i] - col2[i])))
	}
	return result
}

func partTwo(col1, col2 []int) int {
	// create frequency map that shows number of times an int has occurred in col2
	col2Freq := map[int]int{}

	// populate frequency map
	for i := 0; i < len(col2); i++ {
		col2Freq[col2[i]] += 1
	}

	// generate similarity score by multiplying each value in col1 with map
	result := 0
	for i := 0; i < len(col1); i++ {
		result += col1[i] * col2Freq[col1[i]]
	}
	return result

}
