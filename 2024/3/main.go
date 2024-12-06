package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type operand string

const (
	multiply operand = "mul"
	start    operand = "do"
	pause    operand = "don't"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	operations := parseFile(string(f))

	partOneOperations := []string{}
	for _, op := range operations {
		partOneOperations = append(partOneOperations, op[0])
	}
	fmt.Printf("Part 1: %d\n", partOne(partOneOperations))
	fmt.Printf("Part 2: %d\n", partTwo(partOneOperations))
}

func parseFile(file string) [][]string {
	pattern := `(do\(\))|(don't\(\))|(mul\(\d{1,3},\d{1,3}\))`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(file, -1)
	return matches

}

func partOne(operations []string) int {
	result := 0
	for _, operation := range operations {
		operator, values := parseOperation(operation)
		result += performOperation(operator, values)
	}
	return result
}

func partTwo(operations []string) int {
	result := 0
	process := true
	for _, operation := range operations {
		operator, values := parseOperation(operation)
		fmt.Printf("operation: %s, values: %d\n", operation, values)
		if operator == start {
			process = true
		} else if operator == pause {
			process = false
		}

		if process {
			result += performOperation(operator, values)
			fmt.Printf("processing: %s, current result: %d\n", operation, result)
		}
	}

	return result

}

func parseOperation(operation string) (operand, []int) {
	// replace symbols with spaces so that it can be loaded into fields
	re := regexp.MustCompile(`[(),]`)
	operationTrimmed := re.ReplaceAllString(operation, " ")
	operationFields := strings.Fields(operationTrimmed)

	operator := operationFields[0]
	var values []int
	for _, field := range operationFields[1:] {
		value, _ := strconv.Atoi(field)
		values = append(values, value)
	}

	return operand(operator), values
}

func performOperation(operand operand, values []int) int { // perform operation
	var result int
	if operand == multiply {
		result = 1
		for _, value := range values {
			result *= value
		}
	}

	return result

}
