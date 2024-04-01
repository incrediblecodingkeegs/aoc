package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*seed-to-soil map:
50 98 2
52 50 48*/

type Range struct {
	start int
	end   int
	delta int
}

func (r Range) String() string {
	return fmt.Sprintf("start: %d end: %d delta: %d;", r.start, r.end, r.delta)
}

type Seed struct {
	id          int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
	state       SeedState
}

func (s Seed) String() string {
	return fmt.Sprintf("id: %d soil: %d fertilizer: %d water: %d light: %d temperature: %d humidity: %d location: %d state: %d\n",
		s.id, s.soil, s.fertilizer, s.water, s.light, s.temperature, s.humidity, s.location, s.state)
}

type SeedState int

const (
	soil SeedState = iota
	fertilizer
	water
	light
	temperature
	humidity
	location
)

func main() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)
	// Read in seeds from first line in file
	seeds := []*Seed{}
	scanner.Scan()
	seedString := scanner.Text()
	seedStringSplit := strings.Split(seedString, ":")
	seedNumbers := strings.Fields(seedStringSplit[1])

	for _, numberString := range seedNumbers {
		number, _ := strconv.Atoi(numberString)
		seeds = append(seeds, &Seed{
			id: number,
		})
	}

	// Read in next almanac
	scanner.Scan()
	almanacStrings := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			processAlmanac(almanacStrings, seeds)
			almanacStrings = []string{}
		} else {

			almanacStrings = append(almanacStrings, scanner.Text())
		}
	}
	processAlmanac(almanacStrings, seeds)

	// Find lowest location
	m := 0
	for _, seed := range seeds {
		fmt.Printf("Seed %d : %v\n", seed.id, seed)
		if m == 0 {
			m = seed.location
		} else if seed.location < m {
			m = seed.location
		}

	}

	fmt.Printf("Min = %d\n", m)

}

func processAlmanac(almanacStrings []string, seeds []*Seed) {
	// build almanac
	almanac := []*Range{}
	for _, almanacString := range almanacStrings[1:] {
		almanac = append(almanac, getConversionRange(almanacString))
	}

	// update seeds with almanac
	for _, seed := range seeds {
		switch seed.state {
		case soil:
			seed.soil = convertValue(seed.id, almanac)
		case fertilizer:
			seed.fertilizer = convertValue(seed.soil, almanac)
		case water:
			seed.water = convertValue(seed.fertilizer, almanac)
		case light:
			seed.light = convertValue(seed.water, almanac)
		case temperature:
			seed.temperature = convertValue(seed.light, almanac)
		case humidity:
			seed.humidity = convertValue(seed.temperature, almanac)
		case location:
			seed.location = convertValue(seed.humidity, almanac)
		}
		fmt.Printf("almanac = %s\nseed = %v\n", almanac, seed)
		seed.state += 1
	}
}

func getConversionRange(s string) *Range {
	// split string into destination range start, source range start, range length
	split := strings.Fields(s)
	destRangeStart, _ := strconv.Atoi(split[0])
	sourceRangeStart, _ := strconv.Atoi(split[1])
	rangeLength, _ := strconv.Atoi(split[2])

	return &Range{
		start: sourceRangeStart,
		end:   sourceRangeStart + rangeLength - 1,
		delta: destRangeStart - sourceRangeStart,
	}
}

func convertValue(value int, almanac []*Range) int {
	for _, r := range almanac {
		if value >= r.start && value <= r.end {
			return value + r.delta
		}
	}
	return value
}
