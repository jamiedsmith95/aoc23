package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func get_codes(line string) int {
	re := regexp.MustCompile(`[0-9]`)
	matches := re.FindAllString(line, -1)
	if matches == nil {
		return -1
	}
	var code int
	last := len(matches) - 1
	first, err := strconv.Atoi(matches[0])
	if err != nil {
		fmt.Println("ERROR: NOT ABLE TO CONVERT")
		return -1
	}
	last, err = strconv.Atoi(matches[last])
	if err != nil {
		fmt.Println("ERROR: NOT ABLE TO CONVERT")
		return -1
	}
	code = 10*first + last
  fmt.Println(line)
  fmt.Println("code: ", code)
	return code
}

func replace_num(line string) string {
  line = strings.ReplaceAll(line,"zero","zero0zero")
  line = strings.ReplaceAll(line,"one","one1one")
  line = strings.ReplaceAll(line,"two","two2two")
  line = strings.ReplaceAll(line,"three","three3three")
  line = strings.ReplaceAll(line,"four","four4four")
  line = strings.ReplaceAll(line,"five","five5five")
  line = strings.ReplaceAll(line,"six","six6six")
  line = strings.ReplaceAll(line,"seven","seven7seven")
  line = strings.ReplaceAll(line,"eight","eight8eight")
  line = strings.ReplaceAll(line,"nine","nine9nine")
  return line

}


func main() {
	var sum_code int
  filePath := os.Args[1]
  readFile, err := os.Open(filePath)
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  var lines []string
  for fileScanner.Scan() {
    lines = append(lines,fileScanner.Text())
  }
  readFile.Close()
	for _, line := range lines {
    new_line := replace_num(line)

		sum_code += get_codes(string(new_line))
	}

	fmt.Println("sum_code: ", sum_code)

}
