package plugin

import (
	gorobot "github.com/Jel1ySpot/gorobot/pkg/bot"
)

var (
	bot  *gorobot.Bot
	apps []*App
)

func InitialPlugins(initBot *gorobot.Bot) {
	bot = initBot

	for _, app := range apps {
		app.loadEvent.dispatch(bot)
	}

	bot.PrivateMessageEvent.Subscribe(privateMessageHandler)
	bot.GroupMessageEvent.Subscribe(groupMessageHandler)
	bot.TempMessageEvent.Subscribe(tempMessageHandler)
}
