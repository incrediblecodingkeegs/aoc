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

type Almanac struct {
	ranges [7][]*Range
}

func (a *Almanac) String() string {
	return fmt.Sprintf("Almanac:\nsoil: %s\n fertilizer: %s\n water: %s\n light: %s\n temperature: %s\n humidity: %s\n location: %s\n\n",
		a.ranges[0], a.ranges[1], a.ranges[2], a.ranges[3], a.ranges[4], a.ranges[5], a.ranges[6])
}

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
	scanner.Scan()
	seedString := scanner.Text()
	seedStringSplit := strings.Split(seedString, ":")
	seedNumbers := strings.Fields(seedStringSplit[1])

	// Read in almanacs
	almanac := new(Almanac)

	scanner.Scan()
	almanacStrings := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			almanac.addAlmanac(almanacStrings)
			almanacStrings = []string{}
		} else {
			almanacStrings = append(almanacStrings, scanner.Text())
		}
	}
	almanac.addAlmanac(almanacStrings)
	fmt.Println(almanac)

	// Parse seeds one at a time, to find min location
	seedNumber := 0
	m := 0
	for i, numberString := range seedNumbers {
		number, _ := strconv.Atoi(numberString)
		// every second numberString is a range of seeds to be processed
		if i%2 != 0 {
			fmt.Printf("Processing %d seeds from %d\n", number, seedNumber)
			for j := number; j >= 0; j-- {
				seed := &Seed{id: j + seedNumber}
				almanac.processSeed(seed)
				//fmt.Printf("seed after processing: %s\n", seed)
				if m == 0 {
					m = seed.location
				} else if seed.location < m {
					m = seed.location
				}
			}
			fmt.Printf("Current min: %d\n", m)
		} else {
			seedNumber = number
		}
	}

	fmt.Printf("Min = %d\n", m)

}

func (a *Almanac) addAlmanac(strings []string) {
	//fmt.Printf("adding almanac strings: %s\n", strings)
	for i := 0; i < len(a.ranges); i++ {

		if a.ranges[i] == nil {
			a.ranges[i] = []*Range{}
			for _, s := range strings[1:] {
				//fmt.Println(s)
				a.ranges[i] = append(a.ranges[i], getConversionRange(s))
			}
			break
		}
	}
}

func (a *Almanac) processSeed(seed *Seed) {
	for _, ranges := range a.ranges {
		switch seed.state {
		case soil:
			seed.soil = convertValue(seed.id, ranges)
		case fertilizer:
			seed.fertilizer = convertValue(seed.soil, ranges)
		case water:
			seed.water = convertValue(seed.fertilizer, ranges)
		case light:
			seed.light = convertValue(seed.water, ranges)
		case temperature:
			seed.temperature = convertValue(seed.light, ranges)
		case humidity:
			seed.humidity = convertValue(seed.temperature, ranges)
		case location:
			seed.location = convertValue(seed.humidity, ranges)
		}
		//fmt.Printf("almanac = %s\nseed = %v\n", almanac, seed)
		seed.state += 1
	}
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
		//fmt.Printf("almanac = %s\nseed = %v\n", almanac, seed)
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
