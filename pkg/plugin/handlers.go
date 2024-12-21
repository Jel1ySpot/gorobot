package plugin

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strings"
)

func privateMessageHandler(client *client.QQClient, event *message.PrivateMessage) {
	dispatchMessage(fromPrivateMessage(client, event))
}

func groupMessageHandler(client *client.QQClient, event *message.GroupMessage) {
	dispatchMessage(fromGroupMessage(client, event))
}

func tempMessageHandler(client *client.QQClient, event *message.TempMessage) {
	dispatchMessage(fromTempMessage(client, event))
}

func dispatchMessage(msg *MessageContext) {
	for _, app := range apps {
		go app.dispatch(msg)
	}
}

func (a *App) dispatch(ctx *MessageContext) {
	for _, handler := range a.messageHandlers {
		if handler.match(ctx) {
			handler.dispatch(ctx)
		}
	}

	if strings.HasPrefix(ctx.String(), bot.Config.CommandPrefix) {
		cmd, err := buildCommand(ctx)
		if err != nil {
			a.Logger.Debugln("Failed to parse command: ", err)
			return
		}
		for _, handler := range a.commandHandlers {
			if handler.match(cmd) {
				handler.dispatch(cmd)
			}
		}
	}
}
