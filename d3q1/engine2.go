package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func get_gears(input []string) [][][]int {
	re := regexp.MustCompile(`[\*]`)
	var point [][]int
	var points [][][]int

	for _, str := range input {
		point = re.FindAllStringIndex(str, -1)
		if point != nil {
			points = append(points, point)
		} else {

			points = append(points, point)
		}
	}
	return points
}

// get_syms and get gears could be the same function and have it take the regex as input?
func get_syms(input []string) bool {
	re := regexp.MustCompile(`[^\d\.]`)
	var point [][]int
	var found bool
	found = false

	for _, str := range input {
		point = re.FindAllStringIndex(str, -1)
		if point != nil {
			return true
		}
	}
	return found
}

func get_nums(input []string) ([]int, [][][]int) {
	re := regexp.MustCompile(`[\d]+`)
	var nums []int
	var numstr []string
	var numIdx [][][]int
	for _, line := range input {
		numstr = re.FindAllString(line, -1)
		numIdx = append(numIdx, re.FindAllStringIndex(line, -1))

		for _, n := range numstr {
			thisNum, err := strconv.Atoi(n)

			if err != nil {
				fmt.Println(err)
			}
			nums = append(nums, thisNum)
		}

	}

	return nums, numIdx
}

func main() {
	filePath := os.Args[1]
	var sum int
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
	_, numIdx := get_nums(lines)
	gears := get_gears(lines)
	var gear_parts [][]int
	fmt.Println(gears)

	for i, grs := range gears {
		for _, gear := range grs {
			var gear_part []int
			if len(gear) > 0 {
				for k, idxs := range numIdx {
					if len(idxs) > 0 {
						for _, idx := range idxs {

							if (i >= k-1 && i <= k+1) && (gear[0] >= idx[0]-1 && gear[0] <= idx[1]) {
                num, _ := strconv.Atoi(lines[k][idx[0]:idx[1]])
                fmt.Println("idx ",idx)
                fmt.Println("gear ",gear)
                fmt.Println(num)
								gear_part = append(gear_part, num)
							}
						}
					}
				}
			}
			gear_parts = append(gear_parts, gear_part)

		}
	}
	for _, gp := range gear_parts {
    fmt.Println(gp)
		if len(gp) == 2 {
      fmt.Println(gp[0]*gp[1])
			sum = sum + (gp[0] * gp[1])

		}
	}
  fmt.Println(sum)

}
