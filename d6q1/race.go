package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	Time      int
	Distance  int
	Solutions []int
}

var races []Race

func readLine(line string) []string {
	re := regexp.MustCompile(`[\d]+`)
	strs := re.FindAllString(line, -1)
	fmt.Println(strs)
	return strs
}

func (r *Race) performRace(acceleration int) {
	dist := r.Distance
	time := r.Time
	timeLeft := time - acceleration
	if timeLeft*acceleration > dist {
		r.Solutions = append(r.Solutions, acceleration)
	} else {
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
	times := readLine(lines[0])
	dists := readLine(lines[1])
	var longRace Race
	Nraces := len(times)
	var powerTime []int
	var powerTimeTotal int
	var powerDist []int
	var powerDistTotal int
  var timeString string
  var distString string

	for i := 0; i < Nraces; i++ {
		powerTimeTotal += len(times[i])
		powerTime = append(powerTime, powerTimeTotal)
		powerDistTotal += len(dists[i])
		powerDist = append(powerDist, powerDistTotal)
    timeString += times[i] 
    distString += dists[i] 
		time, err := strconv.Atoi(times[i])
		dist, err1 := strconv.Atoi(dists[i])
		if err != nil || err1 != nil {
			fmt.Println(err, err1)
		}

		race := Race{Time: time, Distance: dist}

		races = append(races, race)
	}

  totalTime,err := strconv.Atoi(timeString)
  totalDist,err1 := strconv.Atoi(distString)
  if err != nil || err1 != nil {
    fmt.Println(err,err1)
  }
  longRace = Race {Time: totalTime, Distance: totalDist}
   
	product := 1
	for _, race := range races {
		for i := 0; i < race.Time; i++ {
			race.performRace(i)
		}


		product *= len(race.Solutions)

	}
	var limits []int
	for i := 0; i < longRace.Time; i++ {
		if len(longRace.Solutions) == 0 {
			longRace.performRace(i)
		} else if len(longRace.Solutions) == 1 {
			limits = append(limits, i-1)
			break
		}
	}
	fmt.Println(races[0])
	fmt.Println(longRace)
	for i := longRace.Time; i > limits[0]; i-- {
		if len(longRace.Solutions) == 1 {
			longRace.performRace(i)
		} else if len(longRace.Solutions) == 2 {
			limits = append(limits,i+2)
      break
		}
	}
	fmt.Println(limits)
  fmt.Println(longRace)
  fmt.Println(limits[1]- limits[0])
}
