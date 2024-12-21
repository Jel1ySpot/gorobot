package ping

import "github.com/Jel1ySpot/gorobot/pkg/plugin"

func init() {
	plugin.Create("ping"). // 创建应用
				OnMessage("^ping$").                      // 监听消息
				Handle(func(ctx *plugin.MessageContext) { // 回调函数
			if _, err := ctx.ReplyText("🏓"); err != nil {
				ctx.Logger.Errorln(err)
			}
		})
}
