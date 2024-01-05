package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	seq             []int
	next            *Layer
	extrapolated    int
	preExtrapolated int
}

var layers []Layer

func (l *Layer) getNext() {
	var seqInt []int
	var done bool
	done = true

	for i := 0; i < len(l.seq)-1; i++ {
		diff := l.seq[i+1] - l.seq[i]
		seqInt = append(seqInt, diff)
		if diff != 0 {
			done = false
		}
	}
	l.next = &Layer{seq: seqInt}
	if !done {
		l.next.getNext()
	}
}

func (l *Layer) subFirst() {
	if l.next != nil {
		l.next.subFirst()
		l.preExtrapolated = l.seq[0] - l.next.preExtrapolated
	} else {
		l.preExtrapolated = l.seq[0]
	}
}
func (l *Layer) sumLast() {
	if l.next != nil {
		l.next.sumLast()
		l.extrapolated = l.next.extrapolated + l.seq[len(l.seq)-1]
	} else {
		l.extrapolated = l.seq[len(l.seq)-1]
	}
}

func getLayer(line string) []int {
	var seqInt []int
	strings := strings.Split(line, " ")
	for _, st := range strings {
		valInt, err := strconv.Atoi(st)
		if err != nil {
			fmt.Println(err)
		} else {
			seqInt = append(seqInt, valInt)
		}

	}
	return seqInt
}

func (l *Layer) printLayers() {
	fmt.Println(l.seq)
	if l.next != nil {
		l.next.printLayers()
	}
}

func (l *Layer) extrapolate() int {
	var ex int
	ex = l.seq[0]
	if len(l.seq) == 1 {
		ex = l.seq[0]
	} else if len(l.seq) > 0 && l.next != nil {
		l.next.extrapolate()
		ex = l.seq[len(l.seq)-1] + l.next.extrapolated
	}
	l.extrapolated = ex
	l.seq = append(l.seq, ex)
	return ex
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
	var sum int
	var exps []int
  var preSum int

	for _, l := range lines {
		L := Layer{seq: getLayer(l)}
		L.getNext()
		// L.extrapolate()
		L.printLayers()
		L.sumLast()
    L.subFirst()
		sum += L.extrapolated
    preSum += L.preExtrapolated
		exps = append(exps, L.extrapolated)
		layers = append(layers, L)
	}
	fmt.Println(sum)
	fmt.Println(preSum)
	// fmt.Println(exps)
}
