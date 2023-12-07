package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
  id int
  line string
  game []int
  player []int
  score int
}
type Cards []*Card
var card_list Cards

func (c *Cards) NewCard(line string) Card {
  var card Card
  card.line = line
  err := card.get_details()
  if err != nil {
    fmt.Println(err)
  }
  card_list = append(card_list,&card)
  fmt.Printf("%#v\n", card)
  return card
}
func (c *Card) get_score() {

}


func (c *Card) get_details() error {
  var err error
  idre := regexp.MustCompile(`[\d]+:`)
  numre := regexp.MustCompile(`[\d]+^:`)
  split := strings.Split(c.line, "|")
  c.id,_ = strconv.Atoi(idre.FindString(split[0]))
  game_str := numre.FindAllString(split[0],-1)
  player_str := numre.FindAllString(split[1],-1)
  var player []int
  var game []int
  for _,pm := range player_str {
    player_num,err := strconv.Atoi(pm)
    if err != nil {
      fmt.Println(err)
      return err
    } else {
      player = append(player,player_num)
    }

  }
  for _,gm := range game_str {
    game_num,err := strconv.Atoi(gm)
    if err != nil {
      fmt.Println(err)
      return err
    } else {
      game = append(game,game_num)
    }

  }
  c.player = player
  c.game = game
  return err


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
  for _,line := range lines {
    card_list.NewCard(line)
  }
  // fmt.Printf("%#v",card_list)
  // fmt.Println(card_list[0].game)
}
