package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Seeds struct {
	seeds [][]int
}
type Map struct {
	sourceStart []int
	targetStart []int
	length      []int
}

var maps []Map

func getSeeds(line string) Seeds {
	var seeds [][]int
	var theRange int
	re := regexp.MustCompile(`[\d]+`)
	seedString := re.FindAllString(line, -1)
	for i, s := range seedString {

		seed, err := strconv.Atoi(s)
		var firstSeed []int
		if err != nil {
			fmt.Println(err)
		}
		if i%2 == 0 {
			theRange = seed
		} else {
			for j := 0; j < theRange; j++ {
        seed += j

				firstSeed = append(firstSeed, seed)
				seeds = append(seeds, firstSeed)
			}
		}

	}
	Seed := Seeds{seeds}
	return Seed
}

func (m Map) convert(s *Seeds) {
	if len(m.length) > 0 {
		sourceStart := m.sourceStart
		targetStart := m.targetStart
		length := m.length
		var sourceEnd []int
		for i, strt := range sourceStart {
			sourceEnd = append(sourceEnd, strt+length[i]-1)
		}

		var changed bool
		var thisEnd int

		for j, seed := range s.seeds {
			seedCurrent := seed[len(seed)-1]
			changed = false
			for i, start := range sourceStart {
				thisEnd = sourceEnd[i]
				if seedCurrent >= start && seedCurrent <= thisEnd && !changed {
					seedCurrent = targetStart[i] + (seedCurrent - start)
					changed = true
					seed = append(seed, seedCurrent)
					s.seeds[j] = seed
				} else {

				}
			}
		}
	}
}

func (m *Map) addMap(line string) bool {
	re := regexp.MustCompile(`[\d]+`)
	strs := re.FindAllString(line, -1)
	if len(strs) < 3 {
		return true
	}
	var nums []int
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, n)
	}
	if len(nums) == 3 {
		m.targetStart = append(m.targetStart, nums[0])
		m.sourceStart = append(m.sourceStart, nums[1])
		m.length = append(m.length, nums[2])
		return false
	} else {
		return true
	}
}

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	seeds := getSeeds(lines[0])
	var nextMap Map
	count := 0

	for i, line := range lines[1:] {
		done := nextMap.addMap(line)
		if done || i == len(lines[1:])-1 {
			count++

			maps = append(maps, nextMap)
			nextMap.convert(&seeds)
			nextMap = Map{}
		}

	}
	var final []int
	for _, s := range seeds.seeds {
		theEnd := len(s) - 1
		final = append(final, s[theEnd])
	}
	var lowest int
	lowest = final[0]
	for _, s := range final[1:] {
		if s < lowest {
			lowest = s
		}
	}
	fmt.Println(lowest)

}
