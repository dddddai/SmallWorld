package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const api = "https://ygocdb.com/card/"

var deckPath string
var target string

type card struct {
	id        string
	name      string
	atk       string
	def       string
	level     string
	race      string
	attribute string
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
				deckPath = "./" + f.Name()
				break
			}
		}
	}
	ids := getCardIDs(deckPath)
	//fmt.Println(ids)
	fmt.Println("Monsters:")
	cards := getCards(ids)
	fmt.Println("Routes:")
	getRoutes(cards)
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
		s := string(b)
		start, levelFlag := getlr(s, "] ", "★")
		if levelFlag == -1 {
			// not a monster
			continue
		}
		start += 2
		if err != nil {
			panic("Failed to get card info: " + err.Error())
		}
		c := card{
			id:        id,
			name:      getName(s),
			atk:       getAtk(s[levelFlag:]),
			def:       getDef(s[levelFlag:]),
			level:     getLevel(s[levelFlag:]),
			race:      getRace(s[start:]),
			attribute: getAttribute(s[start:]),
		}
		res.Body.Close()
		cards = append(cards, c)
		fmt.Printf("%+v\n\n", c)
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
						fmt.Printf("%s -> %s -> %s\n", c1.name, c2.name, c3.name)
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

func getRace(s string) string {
	return between(s, "", "/")
}
func getName(s string) string {
	return between(s, "<h2><span>", "</span")
}

func getAttribute(s string) string {
	return between(s, "/", "<br")
}

func getAtk(s string) string {
	return between(s, "] ", "/")
}

func getLevel(s string) string {
	return between(s, "★", "]")
}

func getDef(s string) string {
	return between(s, "/", "<hr")
}

func between(s, left, right string) string {
	l, r := getlr(s, left, right)
	return s[l+len(left) : r]
}

func getlr(s, left, right string) (int, int) {
	l := strings.Index(s, left)
	if l == -1 {
		return -1, -1
	}
	return l, strings.Index(s[l:], right) + l
}
