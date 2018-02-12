package plugin

import (
	"regexp"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Today struct {
	Message string
	Channel string
}

var t Today

func NewToday(message string, channel string) Today {
	t.Message = message
	t.Channel = channel

	return t
}

func (t Today) SendMessage() string {
	var response string

	r := regexp.MustCompile(`^(今日|きょう)は(何|なん)の(日|ひ)$`)
	if r.MatchString(t.Message) {
		doc, err := goquery.NewDocument("https://kids.yahoo.co.jp/today/index.html")
		if err != nil {
			fmt.Print("url scarapping failed")
		}
		res, err := doc.Find("dl#dateDtl > dt > span").Html()
		desc, err := doc.Find("dl#dateDtl > dd").Html()
		if err != nil {
			fmt.Print("dom get failed")
		} else {
			response = "*" + res + "*" + " : " + desc
		}
	}

	return response
}
