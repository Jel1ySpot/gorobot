package message_logger

import "github.com/Jel1ySpot/gorobot/pkg/plugin"

func logger(msg *plugin.MessageContext) {
	log := msg.Logger
	switch msg.Type {
	case plugin.GroupMessage:
		_, name := msg.GroupInfo()
		log.Info("[%s]%s(%d): %s", name, msg.Sender.CardName, msg.Sender.Uin, msg)
	case plugin.PrivateMessage:
		log.Info("%s(%d): %s", msg.Sender.Nickname, msg.Sender.Uin, msg)
	case plugin.TempMessage:
		log.Info("[Temp]%s(%d): %s", msg.Sender.Nickname, msg.Sender.Uin, msg)
	}
}

func init() {
	plugin.Create("message_logger"). // 创建应用
						OnMessage(""). // 监听所有消息
						Handle(logger) // 回调函数
}
