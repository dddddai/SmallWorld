package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const api = "https://ygocdb.com/api/v0/?search="

var deckPath string
var target string
var hand = "展示手牌"
var deck = "展示卡组"
var search = "检索"
var table, _ = gotable.Create(hand, deck, search)

type card struct {
	name      string
	atk       float64
	def       float64
	level     float64
	race      float64
	attribute float64
}

func main() {
	fmt.Println("请输入卡组（.ydk文件）路径（若输入*则自动查找当前目录）：")
	_, err := fmt.Scanln(&deckPath)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("请输入想要检索的卡（若输入*则输出卡组中所有检索路径）：")
	_, err = fmt.Scanln(&target)
	if err != nil {
		panic(err.Error())
	}

	if deckPath == "*" {
		pwd, _ := ioutil.ReadDir(".")
		for _, f := range pwd {
			if strings.HasSuffix(f.Name(), ".ydk") {
				deckPath = f.Name()
				break
			}
		}
	}
	ids := getCardIDs(deckPath)
	fmt.Println("\nMonsters:\n")
	cards := getCards(ids)
	fmt.Println("\nRoutes:")

	table.Align(hand, gotable.Left)
	table.Align(deck, gotable.Left)
	table.Align(search, gotable.Left)

	getRoutes(cards)
	//table.CloseBorder()
	table.PrintTable()
	fmt.Println("按Ctrl+C退出")
	select {}
}

func getCardIDs(deckPath string) []string {
	deck, err := os.Open(deckPath)
	if err != nil {
		panic("Can not find deck file: " + err.Error())
	}
	reader := bufio.NewReader(deck)
	start := false
	m := make(map[string]struct{}, 60)
	ids := make([]string, 0, 60)
	for {
		buf, _, err := reader.ReadLine()
		line := string(buf)
		if err != nil {
			panic("Failed to read deck file: " + err.Error())
		}
		if line == "#main" {
			start = true
		} else if start {
			if line == "#extra" {
				break
			}
			if _, ok := m[line]; !ok {
				ids = append(ids, line)
				m[line] = struct{}{}
			}
		}
	}
	return ids
}

func getCards(ids []string) []card {
	cards := make([]card, 0, len(ids))
	for _, id := range ids {
		res, err := http.Get(api + id)
		if err != nil {
			panic("Failed to get card info: " + err.Error())
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic("Failed to get card info: " + err.Error())
		}
		if !bytes.ContainsRune(b, '★') {
			// not a monster
			continue
		}
		m := make(map[string]interface{})
		err = json.Unmarshal(b[1:len(b)-1], &m)
		if err != nil {
			panic("Failed to get card info: " + err.Error())
		}
		data := m["data"].(map[string]interface{})
		c := card{
			name:      m["cn_name"].(string),
			atk:       data["atk"].(float64),
			def:       data["def"].(float64),
			level:     data["level"].(float64),
			race:      data["race"].(float64),
			attribute: data["attribute"].(float64),
		}
		res.Body.Close()
		cards = append(cards, c)
		fmt.Println(c.name + "\n")
	}
	return cards
}

func getRoutes(cards []card) {
	n := len(cards)
	for i := 0; i < n; i++ {
		c1 := cards[i]
		for j := 0; j < n; j++ {
			c2 := cards[j]
			if check(c1, c2) {
				for k := 0; k < n; k++ {
					c3 := cards[k]
					if check(c2, c3) && (target == "*" || strings.Contains(c3.name, target)) {
						table.AddRow(map[string]string{hand: c1.name, deck: c2.name, search: c3.name})
					}
				}
			}
		}
	}
}

func check(a, b card) bool {
	equal := false
	if a.atk == b.atk {
		equal = true
	}
	if a.def == b.def {
		if equal {
			return false
		}
		equal = true
	}
	if a.level == b.level {
		if equal {
			return false
		}
		equal = true
	}
	if a.attribute == b.attribute {
		if equal {
			return false
		}
		equal = true
	}
	if a.race == b.race {
		if equal {
			return false
		}
		equal = true
	}
	return equal
}
