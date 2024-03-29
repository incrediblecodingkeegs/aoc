package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	dat, _ := os.ReadFile("./input")
	//fmt.Printf("%s", dat)

	sum := sumCalibrationValues(strings.Split(string(dat), "\n"))
	fmt.Printf("sum = %d\n", sum)

}

func sumCalibrationValues(values []string) int {
	if len(values) == 0 {
		return 0
	}

	sum := 0
	for _, v := range values {

		// find first int
		first := 0
		for _, c := range v {
			if unicode.IsDigit(c) {
				first, _ = strconv.Atoi(string(c))
				break
			}
		}

		last := 0
		// find last int
		for i := len(v) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(v[i])) {
				last, _ = strconv.Atoi(string(v[i]))
				break
			}
		}

		sum += (first * 10) + last
	}
	return sum
}
