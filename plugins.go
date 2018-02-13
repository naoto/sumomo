package main

import (
	"./plugin"
)

type Plugins struct {
	Message string
	Channel string
}

var p Plugins

func NewPlugins(message string, channel string) Plugins {
	p.Message = message
	p.Channel = channel

	return p
}

func (t Plugins) Run() []string {
	var res []string
	today := plugin.NewToday(p.Message, p.Channel)
	res = append(res, today.SendMessage())
	res = append(res, plugin.NewWeather(p.Message, p.Channel).SendMessage())

	return res
}
