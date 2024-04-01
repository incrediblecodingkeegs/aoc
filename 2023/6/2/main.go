package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	fSplit := strings.Split(string(f), "\n")
	races := parseRaces(fSplit[0], fSplit[1])

	result := 1
	for _, race := range races {
		result *= getNumWinningChances(race)
	}

	fmt.Printf("result: %d", result)

}

func parseRaces(timeString string, distanceString string) []*Race {
	timeSplit := strings.Split(timeString, ":")
	times := strings.Fields(timeSplit[1])
	distanceSplit := strings.Split(distanceString, ":")
	distances := strings.Fields(distanceSplit[1])

	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(distances, ""))

	races := []*Race{
		&Race{
			time:     t,
			distance: d,
		},
	}

	return races
}

// distance = speed * time
// speed = charge * time
// distance = (charge * t0) * time(total - t0)

// we need to find cases where distance > race distance

// is there a pattern to winning races?  Forms a perfect parabola
// so find threshold, check from opposite end

func getNumWinningChances(race *Race) int {
	// find first winning race
	i := 1
	for ; i < race.time; i++ {
		d := i * (race.time - i)
		if d > race.distance {
			break
		}
	}

	// find last winning race
	j := race.time - i
	for ; j < race.time; j++ {
		d := j * (race.time - j)
		if d <= race.distance {
			break
		}
	}

	// return difference
	//fmt.Printf("i = %d j = %d\n", i, j)
	return j - i
}
