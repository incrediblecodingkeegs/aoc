package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	fmt.Printf("Part One: %d\n", partOne(string(f)))
	fmt.Printf("Part Two: %d\n", partTwo(string(f)))
}

func partOne(input string) int {
	answerSlice, valuesSlice := parseFile(input)

	result := 0
	for i, answer := range answerSlice {
		if partOneEvaluate(answer, valuesSlice[i]) {
			result += answer
		}
	}

	return result
}

func partTwo(input string) int {
	answerSlice, valuesSlice := parseFile(input)

	result := 0
	for i, answer := range answerSlice {
		if partTwoEvaluate(answer, valuesSlice[i]) {
			result += answer
		}
	}

	return result
}

func parseFile(input string) ([]int, [][]int) {
	lines := strings.Split(input, "\n")
	var answer []int
	var values [][]int
	for _, line := range lines {
		lineSplit := strings.Split(line, ":")
		answerValue, _ := strconv.Atoi(lineSplit[0])
		answer = append(answer, answerValue)
		valueString := strings.Fields(lineSplit[1])
		var valueArray []int
		for _, s := range valueString {
			value, _ := strconv.Atoi(s)
			valueArray = append(valueArray, value)
		}
		values = append(values, valueArray)
	}
	return answer, values
}

func partOneEvaluate(answer int, values []int) bool {
	var helper func(current int, remaining []int) bool

	helper = func(current int, remaining []int) bool {
		if len(remaining) == 0 {
			if current == answer {
				return true
			} else {
				return false
			}
		}

		if current*remaining[0] > answer && current+remaining[0] > answer {
			return false
		}
		return helper(current+remaining[0], remaining[1:]) || helper(current*remaining[0], remaining[1:])
	}
	return helper(values[0], values[1:])
}

func partTwoEvaluate(answer int, values []int) bool {
	var helper func(current int, remaining []int) bool

	helper = func(current int, remaining []int) bool {
		if len(remaining) == 0 {
			if current == answer {
				return true
			} else {
				return false
			}
		}

		if concat(current, remaining[0]) > answer && current*remaining[0] > answer && current+remaining[0] > answer {
			return false
		}
		return helper(concat(current, remaining[0]), remaining[1:]) || helper(current+remaining[0], remaining[1:]) || helper(current*remaining[0], remaining[1:])
	}
	return helper(values[0], values[1:])
}

func concat(a, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	concat := aStr + bStr
	res, _ := strconv.Atoi(concat)
	return res
}
