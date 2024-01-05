package main

import (
	"bufio"
	"fmt"
	"os"
)

type step struct {
	current string
	left    string
	right   string
}

type instruction struct {
	steps string
}

func greatestCommonDivisor(a int, b int) int {
  for b!= 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

func calcielateLowestCommonMultiple(a int, b int, integers ...int) int {
  result := a*b/greatestCommonDivisor(a,b)
  for i:=0;i<len(integers);i++ {
    result = calcielateLowestCommonMultiple(result,integers[i])
  }
  return result
}

func getStarts(steps map[string]step) []string {
	var starts []string
	for _, s := range steps {
		if s.current[2] == 'A' {
			starts = append(starts, s.current)
		}

	}
	return starts
}

func isEndZ(s step) bool {
	if s.current[2] == 'Z' {
		return true
	} else {
		return false
	}
}

func doInstruction(ins string, m map[string]step, current string) string {
	c := current
	for _, d := range ins {
		if d == 'L' {
			c = m[string(c)].left
		} else {
			c = m[string(c)].right
		}
	}
	return c
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
	instructions := lines[0]
	var maps map[string]step
	maps = make(map[string]step)
	for _, line := range lines[1:] {
		if len(line) > 3 {
			maps[string(line[:3])] = step{current: string(line[:3]), left: string(line[7:10]), right: string(line[12:15])}

		}
	}
  var stepCount int
	// current := string("AAA")
	// for current != string("ZZZ") {
	// 	for _, d := range instructions {
	// 		stepCount++
	// 		if d == 'L' {
	// 			current = maps[string(current)].left
	// 		} else {
	// 			current = maps[string(current)].right
	// 		}
	// 	}
	// }
	// fmt.Println(stepCount)

	starts := getStarts(maps)

	stepCount = 0.
	notDone := true
	var steps []int
  var orderSteps []int
	for i := 0; i < len(starts); i++ {
		steps = append(steps, 0)
	}

	for i := 0; notDone; i++ {
		notDone = false
		if i == len(instructions) {
			i -= len(instructions)
		}
		instruction := instructions[i]
		stepCount++
		for j, start := range starts {
			starts[j] = doInstruction(string(instruction), maps, start)
			if !isEndZ(maps[string(starts[j])]) {
				notDone = true
			}
			if isEndZ(maps[string(starts[j])]) && steps[j] == 0 {
				steps[j] = stepCount
        orderSteps = append(orderSteps,stepCount)
			}

		}
		done := true
		for _, step := range steps {
			if step == 0 {
				done = false
			}
		}
    if done {
      break
    }
	}
  var lowestCommon int
  lowestCommon = calcielateLowestCommonMultiple(orderSteps[0],orderSteps[1],orderSteps[2:]...)
  fmt.Println(lowestCommon)

  
}
