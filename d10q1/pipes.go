package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type pipe struct {
	pipeType string
	coord    []int
	dist     int
}

type pipes struct {
	startCoord []int
	pipes      []pipe
	valid      bool
	area       int
}

func (p *pipes) getArea(startCoord []int){
  var summation float64
  previousCoord := startCoord
  for _,pipe := range p.pipes {
    currentCoord := pipe.coord
    summation += float64(previousCoord[0]*currentCoord[1]-previousCoord[1]*currentCoord[0])/2
    previousCoord = currentCoord
  }
  enclosed := summation - float64(len(p.pipes)+1)/2 + 1
  fmt.Println(enclosed)
  p.area =int(math.Ceil(enclosed))
}

func (p *pipes) findEnclosedArea(lines []string,startCoord []int) {
	var xMin, yMin, xMax, yMax int
	xMin = len(lines[0])
	yMin = len(lines)
	xMax = 0
	yMax = 0
	var pipeMap map[string]bool
	pipeMap = make(map[string]bool)
  startString := string(startCoord[0]) + " " + string(startCoord[1])
  pipeMap[startString] = true
	for _, pipe := range p.pipes {
		pipeString := string(pipe.coord[0]) + " " + string(pipe.coord[1])
		pipeMap[pipeString] = true
		if pipe.coord[0] < xMin {
			xMin = pipe.coord[0]
		}
		if pipe.coord[0] > xMax {
			xMax = pipe.coord[0]
		}
		if pipe.coord[1] < yMin {
			yMin = pipe.coord[1]
		}
		if pipe.coord[1] > yMax {
			yMax = pipe.coord[1]
		}
	}

	areaCount := 0

	for y := yMin; y <= yMax; y++ {
		var isInside bool
		var upPipe, downPipe bool
		for x := xMin; x <= xMax; x++ {
			currentString := string(x) + " " + string(y)
			if upPipe && downPipe {
				isInside = !isInside
				upPipe = false
				downPipe = false
			}
			if pipeMap[currentString] {
				fmt.Print("-")
				if lines[y][x] == '7' || lines[y][x] == 'F' {
					downPipe = !downPipe
				} else if lines[y][x] == 'J' || lines[y][x] == 'L' {
					upPipe = !upPipe
				} else if lines[y][x] == '|' {
					downPipe = !downPipe
					upPipe = !upPipe
				} else if lines[y][x] == 'S' {
					ydiff := p.pipes[0].coord[1] - p.pipes[len(p.pipes)-1].coord[1]

					if ydiff == 2 || ydiff == -2 {
						downPipe = !downPipe
						upPipe = !upPipe
					} else if int(math.Abs(float64(ydiff))) == 1 && y-p.pipes[0].coord[1] == 1 {
						upPipe = !upPipe

					} else if int(math.Abs(float64(ydiff))) == -1 && y-p.pipes[0].coord[1] == -1 {
						downPipe = !downPipe
					}
				}

			} else if isInside {
				areaCount++
				fmt.Print("0")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
	p.area = areaCount
	fmt.Println(p.area)
}

func (p *pipes) addPipe(lines []string) {
	thisPipe := p.pipes[len(p.pipes)-1]
	var nextX, nextY []int
	x, y := thisPipe.coord[0], thisPipe.coord[1]
	switch thisPipe.pipeType {
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
		nextY = append(nextY, y+1, y)
		break
	case "F":
		nextX = append(nextX, x, x+1)
		nextY = append(nextY, y+1, y)
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
		if len(p.pipes) > 2 {
			p.valid = true
			return
		} else {
			return
		}
	default:
		p.valid = false
		return
	}
	var newCoord []int
	var prevX, prevY int
	if len(p.pipes) > 1 {
		prevX = p.pipes[len(p.pipes)-2].coord[0]
		prevY = p.pipes[len(p.pipes)-2].coord[1]
	} else {
		prevX = p.startCoord[0]
		prevY = p.startCoord[1]
	}
	if nextX[0] < 0 || nextX[1] < 0 || nextY[0] < 0 || nextY[1] < 0 {
		p.valid = false
		return
	}
	if isPipe(lines, nextX[0], nextY[0]) && isPipe(lines, nextX[1], nextY[1]) {
		if prevX == nextX[0] && prevY == nextY[0] {
			newCoord = append(newCoord, nextX[1], nextY[1])
		} else if prevX == nextX[1] && prevY == nextY[1] {
			newCoord = append(newCoord, nextX[0], nextY[0])
		} else {
			p.valid = false
			return
		}
	} else {
		p.valid = false
		return
	}
	newPipe := pipe{pipeType: string(lines[newCoord[1]][newCoord[0]]), coord: newCoord}
	p.pipes = append(p.pipes, newPipe)
	if p.valid {
		p.addPipe(lines)
	}
}

func isPipe(lines []string, x int, y int) bool {
	if x >= 0 && y >= 0 && x < len(lines[0]) && y < len(lines) {

		if lines[y][x] == '.' {
			return false
		} else {
			return true
		}
	} else {
		return false
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

func getStartPipes(lines []string, x int, y int) []pipes {
	var xValid []int
	var yValid []int
	if x > 0 {
		xValid = append(xValid, x-1)
	}
	if x < len(lines[0]) {
		xValid = append(xValid, x+1)
	}
	if y > 0 {
		yValid = append(yValid, y-1)
	}
	if y < len(lines) {
		yValid = append(yValid, y+1)
	}
	var pipeStarts []pipes
	for _, i := range xValid {
		if isPipe(lines, i, y) {
			var newCoord []int
			var startCoord []int
			startCoord = append(startCoord, x, y)
			newCoord = append(newCoord, i, y)
			var newPipes []pipe
			newPipe := pipe{coord: newCoord, pipeType: string(lines[y][i])}
			newPipes = append(newPipes, newPipe)
			newStart := pipes{pipes: newPipes, valid: true, startCoord: startCoord}
			pipeStarts = append(pipeStarts, newStart)
		}
	}

	for _, i := range yValid {
		if isPipe(lines, x, i) {
			var newCoord []int
			var startCoord []int
			startCoord = append(startCoord, x, y)
			newCoord = append(newCoord, x, i)
			var newPipes []pipe
			newPipe := pipe{coord: newCoord, pipeType: string(lines[i][x])}
			newPipes = append(newPipes, newPipe)
			newStart := pipes{pipes: newPipes, valid: true, startCoord: startCoord}
			pipeStarts = append(pipeStarts, newStart)
		}
	}
	return pipeStarts
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
	pipeStarts := getStartPipes(lines, startCoord[0], startCoord[1])
	var fullPipes []pipes
	for _, p := range pipeStarts {
		p.addPipe(lines)
		fullPipes = append(fullPipes, p)
	}

	var validPipes []pipes

	for _, p := range fullPipes {
		if p.valid {
			p.findEnclosedArea(lines,startCoord)
      // p.getArea(startCoord)
			validPipes = append(validPipes, p)
		}
	}
	fmt.Println(validPipes[0].area)
  fmt.Println(validPipes[1].area)

	var currentOne, currentTwo pipe
	for i := 1; true; i++ {
		currentOne = validPipes[0].pipes[i-1]
		currentTwo = validPipes[1].pipes[i-1]
		currentOne.dist = i
		currentTwo.dist = i
		if currentTwo.coord[0] == currentOne.coord[0] && currentOne.coord[1] == currentTwo.coord[1] {
			fmt.Println(currentOne.dist, " ", currentOne.coord)
			break

		}
	}
  fmt.Println(validPipes[0].pipes[0])
  fmt.Println(validPipes[1].pipes[0])
  fmt.Println(startCoord)

}
