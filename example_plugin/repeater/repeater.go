package repeater

import (
	"github.com/Jel1ySpot/gorobot/pkg/plugin"
	"strings"
)

func init() {
	app := plugin.Create("repeater") // 创建应用

	app.OnCommand("repeat").
		Handle(func(ctx *plugin.CommandContext) {
			_, _ = ctx.ReplyText("Start repeating")

			session := app.CreateSession() // 创建 Session

			// 监听消息
			session.OnMessage("").Handle(func(msg *plugin.MessageContext) {
				if msg.Sender.Uin != ctx.Sender.Uin ||
					msg.Type != ctx.Type ||
					strings.Index(msg.String(), "stop_repeat") != -1 {
					return
				}
				if msg.Type == plugin.GroupMessage {
					a, b := msg.GroupInfo()
					c, d := ctx.GroupInfo()
					if a != c || b != d {
						return
					}
				}
				_, _ = ctx.ReplyMessage(msg.Elements)
			})

			// 关闭 Session
			session.OnCommand("stop_repeat").Handle(func(_ *plugin.CommandContext) {
				_, _ = ctx.ReplyText("Stop repeating")
				session.Close()
			})
		})
}
