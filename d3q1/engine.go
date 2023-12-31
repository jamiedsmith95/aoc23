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
      points = append(points,point)
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
	fmt.Println(numIdx)
	fmt.Println(nums)

	return nums, numIdx
}


func main() {
	var sum int
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
	count := 0
	nums, numIdx := get_nums(lines)

	for i, idxs := range numIdx {
		line := lines[i]

		for _, idx := range idxs {
			num := nums[count]
			var xl, xu, yl, yu int
			if i > 0 {
				yl = i - 1
			} else {
				yl = i
			}
			if i < len(lines)-1 {
				yu = i + 2
			} else {
				yu = i + 1
			}
			if idx[0] > 0 {
				xl = idx[0] - 1
			} else {
				xl = idx[0]
			}
			if idx[1] < len(line) {
				xu = idx[1] + 1
			} else {
				xu = idx[1]
			}
			block := lines[yl:yu]
			var input []string
			for _, l := range block {
				input = append(input, l[xl:xu])
			}
			// fmt.Println(input)

			fmt.Println(num, "  --  ", input)
			if get_syms(input) {
				sum += num
			} else {
				fmt.Println(num)
			}
			count += 1

		}

	}
	fmt.Println(sum)

}
