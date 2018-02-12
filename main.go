package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				plg := NewPlugins(ev.Text, ev.Channel)
				res := plg.Run()

				for i := range res {
					if res[i] != "" {
						rtm.SendMessage(rtm.NewOutgoingMessage(res[i], ev.Channel))
					}
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}

}

func main() {
	readConfig()
	api := slack.New(viper.GetString("token"))
	os.Exit(run(api))
}
