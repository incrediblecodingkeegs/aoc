package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type page struct {
	beforeList []int
	afterList  []int
}

func main() {
	f, _ := os.ReadFile("./input.txt")

	fmt.Printf("Part One: %d\n", partOne(string(f)))
	fmt.Printf("Part Two: %d\n", partTwo(string(f)))
}

func partOne(input string) int {
	ruleLines, orderSets := parseFile(input)

	// setup map of rules for page ordering
	rules := newRules(ruleLines)

	result := 0
	for _, set := range orderSets {
		if isCorrectOrder(rules, set) {
			result += set[len(set)/2]
		}
	}

	return result
}

func partTwo(input string) int {
	ruleLines, orderSets := parseFile(input)

	// setup map of rules for page ordering
	rules := newRules(ruleLines)

	result := 0
	for _, set := range orderSets {
		if !isCorrectOrder(rules, set) {
			set = makeCorrectOrder(rules, set)
			result += set[len(set)/2]
		}
	}

	return result
}

func parseFile(input string) ([]string, [][]int) {
	// parse input to split out lines that define page order rules and current page order
	lines := strings.Split(input, "\n")
	var ruleLines, orderLines []string
	for _, line := range lines {
		if strings.Contains(line, "|") {
			ruleLines = append(ruleLines, line)
		} else if strings.Contains(line, ",") {
			orderLines = append(orderLines, line)
		}
	}

	// parse sets of orders
	orderSets := make([][]int, len(orderLines))
	for i, line := range orderLines {
		lineSplit := strings.Split(line, ",")
		orderSets[i] = make([]int, len(lineSplit))
		for j, s := range lineSplit {
			orderSets[i][j], _ = strconv.Atoi(s)
		}
	}

	return ruleLines, orderSets
}

func newRules(input []string) map[int]*page {
	rules := make(map[int]*page)
	for _, line := range input {
		parts := strings.Split(line, "|")
		page1, _ := strconv.Atoi(parts[0])
		page2, _ := strconv.Atoi(parts[1])

		if _, ok := rules[page1]; ok {
			rules[page1].afterList = append(rules[page1].afterList, page2)
		} else {
			rules[page1] = &page{afterList: []int{page2}}
		}

		if _, ok := rules[page2]; ok {
			rules[page2].beforeList = append(rules[page2].beforeList, page1)
		} else {
			rules[page2] = &page{beforeList: []int{page1}}
		}
	}

	return rules
}

func isCorrectOrder(rules map[int]*page, order []int) bool {
	for i, currentPage := range order {
		// for each page in order, check for existing rules regarding other pages in the list
		for j := 0; j < len(order); j++ {
			if j == i { // do not check current page
				continue
			}
			pageToFind := order[j]
			// if page being checked comes before the current page in the order
			if j < i {
				// check that page is not meant to come after current page
				if arrayContains(rules[currentPage].afterList, pageToFind) {
					return false
				}
			} else { // page being checked comes after the current page in the order
				// check that page is not meant to come before current page
				if arrayContains(rules[currentPage].beforeList, pageToFind) {
					return false
				}
			}

		}
	}

	return true
}

func arrayContains(arr []int, n int) bool {
	for _, a := range arr {
		if a == n {
			return true
		}
	}
	return false
}

// Scan through and find last number that it goes after and place it there
// [5, 4, 1, 3, 2]	[75]
// [4, 1, 3, 2, 5]	[97,75]
// [1, 3, 2, 4, 5]	[97,75,47]
// [1, 3, 2, 4, 5]	[97,75,47,61]
// [1, 2, 3, 4, 5]	[97,75,47,61,53]
func makeCorrectOrder(rules map[int]*page, order []int) []int {
	result := []int{order[0]}
	lastFoundIndex := 0
	for i := 1; i < len(order); i++ {
		lastFoundIndex = -1
		for j := 0; j < len(result); j++ {
			if arrayContains(rules[order[i]].beforeList, result[j]) {
				lastFoundIndex = j
			}
		}

		// insert page into result array
		if lastFoundIndex != -1 {
			if lastFoundIndex == len(result)-1 { // place page at the end
				result = append(result, order[i])
			} else { // place page in the middle
				result = slices.Concat(result[:lastFoundIndex+1], []int{order[i]}, result[lastFoundIndex+1:])
			}
		} else { // place page at the start
			result = slices.Concat([]int{order[i]}, result)
		}
	}

	return result
}
