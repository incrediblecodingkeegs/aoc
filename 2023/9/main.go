package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPreviousValue(history []int) int {
	predictions := buildPredictions(history)

	// calculate previous value for each row
	for i := len(predictions) - 1; i >= 1; i-- {
		previous := predictions[i-1][0] - predictions[i][0]

		// prepend previous value to previous row
		predictions[i-1] = append([]int{previous}, predictions[i-1]...)
	}

	//fmt.Printf("predictions: %+v\n", predictions)

	return predictions[0][0]

}

func buildPredictions(history []int) [][]int {
	predictions := [][]int{history}
	isFound := false

	// build list of next predictions
	for i := 0; !isFound; i++ {
		isFound = true
		predictions = append(predictions, make([]int, 0))
		for j := 0; j < len(predictions[i])-1; j++ {
			// find  difference between two values
			diff := predictions[i][j+1] - predictions[i][j]

			// not final row if row contains any differences > 0
			if diff != 0 {
				isFound = false
			}

			// add value to next row of predictions
			predictions[i+1] = append(predictions[i+1], diff)
		}
	}

	return predictions
}

func getNextValue(history []int) int {
	predictions := buildPredictions(history)

	//fmt.Printf("predictions after: %v\n", predictions)

	// calculate next value for each row
	for i := len(predictions) - 1; i >= 1; i-- {
		nextValue := predictions[i][len(predictions[i])-1] + predictions[i-1][len(predictions[i-1])-1]

		// append next value to previous row
		predictions[i-1] = append(predictions[i-1], nextValue)
	}

	//fmt.Printf("predictions: %+v\n", predictions)

	return predictions[0][len(predictions[0])-1]
}

func parseHistoriesInput(input string) [][]int {
	split := strings.Split(input, "\n")
	histories := make([][]int, 0)

	// parse all ints in input file
	for _, s := range split {
		history := make([]int, 0)

		// convert string to slice of int strings
		fields := strings.Fields(s)
		for _, field := range fields {

			// convert to int
			value, _ := strconv.Atoi(field)
			history = append(history, value)
		}
		histories = append(histories, history)
	}

	return histories
}

func part2(input string) int {
	histories := parseHistoriesInput(input)

	//fmt.Printf("histories after parsing: %v\n", histories)

	result := 0
	for _, history := range histories {
		v := getPreviousValue(history)
		result += v
		//fmt.Printf("nextValue: %d\n", v)
	}

	return result
}

func part1(input string) int {
	histories := parseHistoriesInput(input)

	//fmt.Printf("histories after parsing: %v\n", histories)

	result := 0
	for _, history := range histories {
		v := getNextValue(history)
		result += v
		//fmt.Printf("nextValue: %d\n", v)
	}

	return result
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to open input file: %v\n", err))
	}
	fmt.Printf("part1: %d\n", part1(string(file)))
	fmt.Printf("part2: %d\n", part2(string(file)))
}
