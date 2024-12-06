package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")

	var reports [][]int
	for _, line := range lines {
		fields := strings.Fields(line)
		report := make([]int, len(fields))
		for i, field := range fields {
			report[i], _ = strconv.Atoi(field)
		}
		reports = append(reports, report)
	}

	result := partOne(reports)
	log.Printf("Part 1: %d\n", result)

	result = partTwo(reports)
	log.Printf("Part 2: %d\n", result)

}

func partOne(reports [][]int) int {
	result := 0
	for _, report := range reports {
		if isSafe(report) {
			result += 1
			//fmt.Printf("%d - safe\n", report)
		} else {
			//fmt.Printf("%d - unsafe\n", report)
		}
	}

	return result
}

func canBeMadeSafe(report []int) bool {
	// attempt to delete each index individually
	for i := 0; i < len(report); i++ {
		tmpReport := make([]int, len(report))
		copy(tmpReport, report)
		slices.Delete(tmpReport, i, i+1)
		tmpReport = tmpReport[:len(tmpReport)-1]
		if isSafe(tmpReport) {
			return true
		}
	}
	return false
}

func isSafe(report []int) bool {
	safe := true

	for i := 1; i < len(report); i++ {
		diff := abs(report[i] - report[i-1])

		if diff > 3 || diff < 1 {
			safe = false
		}
	}

	if !isStrictlyIncreasing(report) && !isStrictlyDecreasing(report) {
		safe = false
	}

	return safe
}

func isStrictlyIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] >= report[i] {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] >= report[i-1] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func partTwo(reports [][]int) int {
	result := 0
	for _, report := range reports {
		if isSafe(report) {
			result += 1
			//fmt.Printf("%d - safe\n", report)
		} else if canBeMadeSafe(report) {
			result += 1
			//fmt.Printf("%d - can be made safe\n", report)
		} else {
			//fmt.Printf("%d - unsafe\n", report)
		}
	}

	return result
}

//
//func partTwo(reports [][]int) int {
//	result := 0
//	for _, report := range reports {
//		leftIndex, rightIndex := 0, len(report)-1
//		left := report[leftIndex]
//		right := report[rightIndex]
//		dropped, safe := false, true
//		dir := getDirection(left, right)
//
//		fmt.Printf("%d ****\n", report[leftIndex:rightIndex+1])
//
//		for leftIndex < rightIndex {
//			left = report[leftIndex]
//			right = report[rightIndex]
//			isTargetPossible := checkTargetPossible(left, right, rightIndex-leftIndex, dir)
//			isNextValuePossible := checkTargetPossible(left, report[leftIndex+1], 1, dir)
//			if isTargetPossible && isNextValuePossible {
//				fmt.Printf("%d %v\n", report[leftIndex:rightIndex+1], isTargetPossible)
//				leftIndex++
//			} else if !dropped { // if a number has not yet been dropped from this report, see if there is another path
//
//				// if right number is only one remaining and we haven't dropped, drop it.
//				if leftIndex == rightIndex-1 {
//					// drop the right number
//					fmt.Printf("%d true -- drop right\n", report[leftIndex:rightIndex+1])
//					dropped = true
//					rightIndex--
//					continue
//				}
//
//				// in some cases it may be the next value that needs to be dropped
//				if !isNextValuePossible {
//					// next value can be dropped if current val can move to n + 2 index without issue
//					if checkTargetPossible(left, report[leftIndex+2], 1, dir) {
//						// drop the next number
//						fmt.Printf("%d %v -- dropping next value\n", report[leftIndex:rightIndex+1], isTargetPossible)
//						dropped = true
//						leftIndex += 2
//						continue
//					}
//				}
//
//				// check whether dir needs to change (in 10 5 6 7 8 edge case)
//				if leftIndex == 0 {
//					dir = getDirection(report[leftIndex+1], right)
//					fmt.Printf("changing dir to %v\n", dir)
//				}
//
//				// check if dropping left number works
//				isTargetPossible = true
//				if leftIndex > 0 {
//					// first check path from previous number if it exists
//					left = report[leftIndex-1]
//					isTargetPossible = checkTargetPossible(left, right, rightIndex-leftIndex, dir)
//
//					// check that the previous number in report will still work with next number
//					isNextValuePossible = checkTargetPossible(left, report[leftIndex+1], 1, dir)
//					isTargetPossible = isTargetPossible && isNextValuePossible
//				}
//				if isTargetPossible && leftIndex < rightIndex-1 {
//					// check path from next number in sequence
//					left = report[leftIndex+1]
//					isTargetPossible = checkTargetPossible(left, right, rightIndex-leftIndex-1, dir)
//					if isTargetPossible {
//						// drop the left number
//						fmt.Printf("%d %v -- drop left\n", report[leftIndex:rightIndex+1], isTargetPossible)
//						dropped = true
//						leftIndex++
//						continue
//					}
//				}
//
//				// check if dropping right number works if dropping the left target did not work
//				if leftIndex < rightIndex-1 {
//					// check path from next number in sequence
//					left = report[leftIndex]
//					right = report[rightIndex-1]
//					isTargetPossible = checkTargetPossible(left, right, rightIndex-1-leftIndex, dir)
//					if isTargetPossible {
//						// drop the right number
//						fmt.Printf("%d %v -- drop right\n", report[leftIndex:rightIndex+1], isTargetPossible)
//						dropped = true
//						rightIndex--
//						continue
//					}
//				}
//
//				if !dropped {
//					// else not possible, set safe to false and break
//					fmt.Printf("%d %v -- cant drop\n", report[leftIndex:rightIndex+1], isTargetPossible)
//					safe = false
//					break
//				}
//
//			} else {
//				fmt.Printf("%d %v %v -- already dropped\n", report[leftIndex:rightIndex+1], isTargetPossible, isNextValuePossible)
//				safe = false
//				break
//			}
//		}
//
//		if safe {
//			result += 1
//		}
//	}
//	return result
//}
