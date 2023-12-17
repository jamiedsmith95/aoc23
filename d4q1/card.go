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
  matches int
  copies int
}
type Cards []*Card
var card_list Cards

func (c *Cards) NewCard(line string) Card {
  var card Card
  card.line = line
  card.copies = 1
  err := card.get_details()
  if err != nil {
    fmt.Println(err)
  }
  card_list = append(card_list,&card)
  return card
}
func (c *Card) get_score() {
  var score int
  var matches int
  score = 0
  player := c.player
  game := c.game
  for _,i := range game {
    for _,j := range player {
      if i == j && score == 0 {
        matches = 1
        score = 1
      } else if i == j {
        score = score *2
        matches += 1
      } else {
      }
    }
  }
  c.score = score
  c.matches = matches

  }


func (c *Card) get_details() error {
  var err error
  idre := regexp.MustCompile(`[\d]+`)
  numre := regexp.MustCompile(`[\d]+`)
  split := strings.SplitN(c.line, "|",2)

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
  for _,gm := range game_str[1:] {
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
  c.get_score()
  return err


}

func main() {
  var sum int
  var totalCards int
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
  for i,s := range card_list {
    N := s.matches
    for j:= 1;j<=N && j+i < len(card_list);j++ {
      card_list[i+j].copies += s.copies
    }
    // fmt.Println(s)


    sum += s.score
    totalCards += s.copies
  }
  fmt.Printf("%#v",card_list[len(card_list)-3])
  fmt.Println(totalCards)
  fmt.Println(sum)
}

