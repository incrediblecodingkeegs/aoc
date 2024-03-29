package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	dat, err := os.ReadFile("./input")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	sum := 0
	for _, s := range strings.Split(string(dat), "\n") {
		v := getCalibrationValue(s)
		sum += v
	}
	fmt.Printf("Sum = %d\n", sum)
	fmt.Printf("Elapsed time : %s\n", time.Since(start))
}

func getCalibrationValue(value string) int {
	re, _ := regexp.Compile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	// find strings of all matches in string
	indexes := re.FindAllStringIndex(value, -1)
	strings := re.FindAllString(value, -1)
	if len(strings) == 0 {
		return 0
	}
	first, err := parseValue(strings[0])
	if err != nil {
		//fmt.Printf("Error: %v\n", err)
		return 0
	}

	// check if overlapping number in second string
	var second = 0
	secondStrings := re.FindAllString(value[indexes[len(indexes)-1][1]-1:], -1)
	if len(secondStrings) != 0 {
		second, err = parseValue(secondStrings[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return 0
		}
	} else {

		second, err = parseValue(strings[len(strings)-1])

	}
	//fmt.Printf("value = %s, first = %d, second = %d, indexes = %v\n", value, first, second, indexes)
	return (first * 10) + second
}

func parseValue(s string) (int, error) {
	stringValues := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	if unicode.IsDigit(rune(s[0])) {
		return strconv.Atoi(string(s[0]))
	}
	if v, ok := stringValues[s]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("could not parse value %s\n", s)
}
