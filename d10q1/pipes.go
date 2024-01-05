package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pipe struct {
	pipeType string
	coord    []int
	dist     int
}

type pipes struct {
	pipes []pipe
  valid bool
}

func (p *pipes) addPipe(lines []string) {
	pipe := p.pipes[len(p.pipes)-1]
	var nextX, nextY []int
	x, y := pipe.coord[0], pipe.coord[1]
	switch pipe.pipeType {
	case "L":
		nextX = append(nextX, x, x+1)
		nextY = append(nextY, y-1, y)
		break
	case "J":
		nextX = append(nextX, x, x-1)
		nextY = append(nextY, y-1, y)
		break
	case "7":
		nextX = append(nextX, x, x-1)
		nextY = append(nextY, y-1, y)
		break
	case "F":
		nextX = append(nextX, x, x+1)
		nextY = append(nextY, y-1, y)
		break
	case "|":
		nextX = append(nextX, x, x)
		nextY = append(nextY, y-1, y+1)
		break
	case "-":
		nextX = append(nextX, x-1, x+1)
		nextY = append(nextY, y, y)
		break
  case "S":
    if len(p.pipes) >2 {
      return 
    }
	default:
		fmt.Println("not pipe, check isPipe()")
	}




}

func isPipe(lines []string, x int, y int) bool {
	if lines[y][x] == '.' {
		return false
	} else {
		return true
	}
}

func getStart(lines []string) []int {
	var startCoords []int
	for i, line := range lines {
		length := strings.Split(line, "S")
		if len(length) > 1 {
			startCoords = append(startCoords, len(length[0]), i)
			return startCoords
		}
	}
	fmt.Println("NO START FOUND")
	return nil
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
	startCoord := getStart(lines)
	fmt.Println(startCoord)
}
