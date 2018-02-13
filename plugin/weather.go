package plugin

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Weather struct {
	Message string
	Channel string
}

func NewWeather(message string, channel string) *Weather {
	var t = new(Weather)
	t.Message = message
	t.Channel = channel

	return t
}

func (t Weather) SendMessage() string {
	var response string

	r := regexp.MustCompile(`^(今日|きょう|明日|あした|あす|明後日|あさって)の(天気|てんき)$`)
	if r.MatchString(t.Message) {
		rr := r.FindAllStringSubmatch(t.Message, -1)
		day := rr[0][1]

		doc, err := goquery.NewDocument("http://www.jma.go.jp/jp/week/353.html")
		if err != nil {
			fmt.Print("url scrapping failed")
		}

		var res string
		var tenki string
		var max string
		var min string
		if day == "今日" || day == "きょう" {
			res = doc.Find("#infotablefont > tbody > tr:nth-child(1) > th:nth-child(2)").Text()
			tenki = doc.Find("#infotablefont > tbody > tr:nth-child(4) > td:nth-child(2)").Text()
			max = doc.Find("#infotablefont > tbody > tr:nth-child(7) > td:nth-child(3)").Text()
			min = doc.Find("#infotablefont > tbody > tr:nth-child(8) > td:nth-child(2)").Text()
		} else if day == "明日" || day == "あす" || day == "あした" {
			res = doc.Find("#infotablefont > tbody > tr:nth-child(1) > th:nth-child(3)").Text()
			tenki = doc.Find("#infotablefont > tbody > tr:nth-child(4) > td:nth-child(3)").Text()
			max = doc.Find("#infotablefont > tbody > tr:nth-child(7) > td:nth-child(4)").Text()
			min = doc.Find("#infotablefont > tbody > tr:nth-child(8) > td:nth-child(3)").Text()
		} else if day == "明後日" || day == "あさって" {
			res = doc.Find("#infotablefont > tbody > tr:nth-child(1) > th:nth-child(4)").Text()
			tenki = doc.Find("#infotablefont > tbody > tr:nth-child(4) > td:nth-child(4)").Text()
			max = doc.Find("#infotablefont > tbody > tr:nth-child(7) > td:nth-child(5)").Text()
			min = doc.Find("#infotablefont > tbody > tr:nth-child(8) > td:nth-child(4)").Text()
		}

		max = strings.Replace(max, "\n", "", -1)
		max = strings.Replace(max, "\t", "", -1)
		min = strings.Replace(min, "\n", "", -1)
		min = strings.Replace(min, "\t", "", -1)

		response = "沖縄の天気: " + res + "\n" + tenki + "\n最高気温:" + max + "\n最低気温:" + min
	}
	return response
}
