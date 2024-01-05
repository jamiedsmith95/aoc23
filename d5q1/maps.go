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
	next  [][]int
}
type Map struct {
	sourceStart []int
	targetStart []int
	length      []int
}

func getOverlap(seedRange []int, mapRange []int) [][]int {
	var returnRanges [][]int
	if len(seedRange) < 1 {
		returnRanges = append(returnRanges, seedRange)
		fmt.Println("skip")
		return nil
	}
	var Range []int
	var rnge int
	if (seedRange[0] >= mapRange[0] && seedRange[0] <= mapRange[1]) || (mapRange[0] >= seedRange[0] && mapRange[0] <= seedRange[1]) {
		var start, end int
		if seedRange[0] >= mapRange[0] {
			rnge = min(seedRange[1], mapRange[1]) - seedRange[0]
			start = mapRange[2] + seedRange[0] - mapRange[0]
			end = start + rnge
		} else if mapRange[0] > seedRange[0] {
			start = mapRange[2]
			rnge = min(seedRange[1], mapRange[1]) - mapRange[0]
			end = start + rnge
		}
		Range = append(Range, start, end)
		returnRanges = append(returnRanges, Range)
		if seedRange[0] < mapRange[0] {
			var preRange []int
			preRange = append(preRange, seedRange[0])
			preRange = append(preRange, mapRange[0]-1)
			returnRanges = append(returnRanges, preRange)
		}
		if seedRange[1] > mapRange[1] {
			var postRange []int
			postRange = append(postRange, mapRange[1]+1)
			postRange = append(postRange, seedRange[1])
			returnRanges = append(returnRanges, postRange)
		}
	} else {
		return nil
	}
	return returnRanges
}

func (s *Seeds) mapRanges(m []int) {
	for i, seedR := range s.seeds {
		if seedR != nil {
			returnRanges := getOverlap(seedR, m)
			count += len(returnRanges) 
			if returnRanges != nil  {
				s.seeds[i] = nil
        s.next = append(s.next,returnRanges[0])
        count -=1
				// fmt.Println("comparison")

        for _, ret := range returnRanges[1:] {

					s.seeds = append(s.seeds, ret)
				}
			}
		}
	}
}

func getRangeSeeds(line string) Seeds {
	var seeds [][]int
	var theRange int
	re := regexp.MustCompile(`[\d]+`)
	seedString := re.FindAllString(line, -1)
	var firstSeed int
	for i, s := range seedString {

		seed, err := strconv.Atoi(s)
		var seedRange []int
		if err != nil {
			fmt.Println(err)
		}
		if i%2 == 0 {
			firstSeed = seed
		} else {
			theRange = seed
			seedRange = append(seedRange, firstSeed, firstSeed+theRange-1)
			seeds = append(seeds, seedRange)
		}

	}
	Seed := Seeds{seeds, nil}
	return Seed
}

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
	Seed := Seeds{seeds, nil}
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
					// seed = append(seed, seedCurrent)
					seed[0] = seedCurrent
					s.seeds[j] = seed
				} else {

				}
			}
		}
	}
}

var count int

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
	if len(nums) > 1 {
		m.targetStart = append(m.targetStart, nums[0])
		m.sourceStart = append(m.sourceStart, nums[1])
		m.length = append(m.length, nums[2])
		return false

	} else {
		return true
	}
}
func (s *Seeds) getLowest() int {
	var lowest int
	var mark bool
	mark = false
	for _, s := range s.seeds {
		if len(s) > 0 {
			if !mark || s[0] < lowest && s[0] > 0 {
				fmt.Println(s)
				lowest = s[0]
				mark = true
			}
		}

	}
  return lowest

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
	seeds := getRangeSeeds(lines[0])

	var nextMap Map

	for i, line := range lines[1:] {
		done := nextMap.addMap(line)
		if done && len(nextMap.length) > 1 || i == len(lines[1:])-1 {
			for j, l := range nextMap.length {
				var mapping []int
				mapping = append(mapping, nextMap.sourceStart[j], nextMap.sourceStart[j]+l-1, nextMap.targetStart[j])
				seeds.mapRanges(mapping)

			}

			nextMap = Map{}
      for _,s := range seeds.seeds {
        if s != nil {
          seeds.next = append(seeds.next, s)
        }

      }
			seeds.seeds = seeds.next

			fmt.Println("the lowest of seeds", seeds.getLowest())
      fmt.Println("the length of seeds", len(seeds.seeds))
      fmt.Println(seeds.seeds)
			seeds.next = nil

		}
	}
	fmt.Println("###########")

	fmt.Println("~######i########################################################################")
	fmt.Println("len of seeds: ", len(seeds.seeds))
	fmt.Println("count: ", count)
	fmt.Println(seeds.getLowest())

}
