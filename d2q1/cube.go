package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func get_id(line string) int {
	re := regexp.MustCompile(`Game\s([\d]+)`)
	match := re.FindAllStringSubmatch(line, -1)
	id, err := strconv.Atoi(match[0][1])
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func get_rgb(line string) []int {
	red := regexp.MustCompile(`([\d]+) red`)
	green := regexp.MustCompile(`([\d]+) green`)
	blue := regexp.MustCompile(`([\d]+) blue`)

	rmatch := red.FindAllStringSubmatch(line, -1)
	gmatch := green.FindAllStringSubmatch(line, -1)
	bmatch := blue.FindAllStringSubmatch(line, -1)

	var rmax, gmax, bmax int
	for _, r := range rmatch {
		R, err := strconv.Atoi(string(r[1]))
		if err != nil {
			fmt.Println(err)
		}
		if R > rmax {
			rmax = R
		}
	}
	for _, b := range bmatch {
		B, err := strconv.Atoi(string(b[1]))
		if err != nil {
			fmt.Println(err)
		}
		if B > bmax {
			bmax = B
		}
	}
	for _, g := range gmatch {
		G, err := strconv.Atoi(string(g[1]))
		if err != nil {
			fmt.Println(err)
		}
		if G > gmax {
			gmax = G
		}
	}
	var rgb []int
	rgb = append(rgb, rmax, gmax, bmax)
	return rgb
}

func possible(line string, tot []int) bool {
	rgb := get_rgb(line)
	if tot[0] < rgb[0] || tot[1] < rgb[1] || tot[2] < rgb[2] {
		fmt.Println(tot, "less than ", rgb)
		return false
	}
	return true
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

	var rgb []int
	for _, line := range lines {
		rgb = get_rgb(line)
		sum += rgb[0] * rgb[1] * rgb[2]
	}
	fmt.Println(sum)

}
