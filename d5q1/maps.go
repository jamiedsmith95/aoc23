package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Seeds struct {
	seeds []int
}
type Map struct {
	sourceStart []int
	targetStart []int
	length      []int
}


var maps []Map

func getSeeds(line string) Seeds {
	var seeds []int
	re := regexp.MustCompile(`[\d]+`)
	seedString := re.FindAllString(line,-1)
	for _, s := range seedString {
		seed, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		seeds = append(seeds, seed)

	}
	Seed := Seeds{seeds}
	return Seed

}


func (m *Map) addMap(line string) bool {
  re := regexp.MustCompile(`[\d]+`)
  strs := re.FindAllString(line,-1)
  var nums []int
  for _,s := range strs {
    n,err := strconv.Atoi(s)
    if err != nil {
      fmt.Println(err)
    }
    nums = append(nums,n)
  }
  if len(nums) == 3 {
    m.targetStart  = append(m.targetStart,nums[0])
    m.sourceStart = append(m.sourceStart,nums[1])
    m.length = append(m.length,nums[2])
    fmt.Println(nums)
    return false
  } else {
    fmt.Println(nums)
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
	// seed := getSeeds(lines[0])
  var nextMap Map

  for _,line := range lines[1:] {
    done := nextMap.addMap(line)
    if done {
      maps = append(maps,nextMap)
      nextMap = Map {}
    }

  }
  fmt.Println(maps)


}
