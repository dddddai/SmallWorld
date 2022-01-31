package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gen2brain/dlgs"
	"github.com/liushuochen/gotable"
)

const api = "https://ygocdb.com/api/v0/?search="

var target = "*"

const (
	hand   = "展示手牌"
	deck   = "展示卡组"
	search = "检索"
)

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
	deckPath, ok, _ := dlgs.File("请选择要查询的卡组文件", "*.ydk", false)
	if !ok {
		dlgs.Error("Error", "Faied to open deck file")
		return
	}
	target, ok, _ = dlgs.Entry("请输入要检索的卡名", "要检索的卡(输入*表示检索所有怪兽)", "")
	if !ok {
		dlgs.Error("Error", "Faied to read target")
		return
	}
	ids := getCardIDs(deckPath)
	fmt.Println("Monsters:\n")
	cards := getCards(ids)
	fmt.Println("\nRoutes:")

	table.Align(hand, gotable.Left)
	table.Align(deck, gotable.Left)
	table.Align(search, gotable.Left)

	getRoutes(cards)
	// table.CloseBorder()
	table.PrintTable()
	fmt.Println("\n查询完毕，感谢使用^_^\n")
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
		res.Body.Close()
		if err != nil {
			panic("Failed to get card info: " + err.Error())
		}

		m := make(map[string]interface{})
		err = json.Unmarshal(b, &m)
		if err != nil {
			panic("Failed to get card info: " + err.Error())
		}

		info := m["result"].([]interface{})[0].(map[string]interface{})

		if !strings.HasPrefix(info["text"].(map[string]interface{})["types"].(string), "[怪") {
			// not a monster
			continue
		}

		data := info["data"].(map[string]interface{})

		c := card{
			name:      info["cn_name"].(string),
			atk:       data["atk"].(float64),
			def:       data["def"].(float64),
			level:     data["level"].(float64),
			race:      data["race"].(float64),
			attribute: data["attribute"].(float64),
		}
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
