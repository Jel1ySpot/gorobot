package plugin

import (
	"fmt"
	gorobot "github.com/Jel1ySpot/gorobot/pkg/bot"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"regexp"
)

type MessageType int

const (
	PrivateMessage MessageType = iota
	TempMessage
	GroupMessage
)

type MessageContext struct {
	*client.QQClient
	Type     MessageType
	Sender   *message.Sender
	Elements []message.IMessageElement
	Logger   gorobot.Logger

	privateMsg *message.PrivateMessage
	tempMsg    *message.TempMessage
	groupMsg   *message.GroupMessage
}

func (m *MessageContext) String() string {
	return message.ToReadableString(m.Elements)
}

func (m *MessageContext) GroupInfo() (uin uint32, name string) {
	if m.groupMsg != nil {
		uin = m.groupMsg.GroupUin
		name = m.groupMsg.GroupName
	}
	return
}

func buildMessageHandle(pattern string) *EventHandle[*MessageContext] {
	var matcher func(ctx *MessageContext) bool
	if pattern != "" {
		re := regexp.MustCompile(pattern)
		matcher = func(ctx *MessageContext) bool {
			return re.FindString(ctx.String()) != ""
		}
	}
	return &EventHandle[*MessageContext]{
		matcher: matcher,
	}
}

func fromPrivateMessage(client *client.QQClient, event *message.PrivateMessage) *MessageContext {
	return &MessageContext{
		QQClient:   client,
		Type:       PrivateMessage,
		Sender:     event.Sender,
		Elements:   event.Elements,
		privateMsg: event,
	}
}

func fromTempMessage(client *client.QQClient, event *message.TempMessage) *MessageContext {
	return &MessageContext{
		QQClient: client,
		Type:     TempMessage,
		Sender:   event.Sender,
		Elements: event.Elements,
		tempMsg:  event,
	}
}

func fromGroupMessage(client *client.QQClient, event *message.GroupMessage) *MessageContext {
	return &MessageContext{
		QQClient: client,
		Type:     GroupMessage,
		Sender:   event.Sender,
		Elements: event.Elements,
		groupMsg: event,
	}
}

func (m *MessageContext) ReplyMessage(content []message.IMessageElement, quote ...bool) (*MessageContext, error) {
	if quote != nil && quote[0] {
		var (
			groupUin uint32
			id       uint32
			time     uint32
		)
		switch m.Type {
		case PrivateMessage:
			id = m.privateMsg.ID
			time = m.privateMsg.Time
		case TempMessage:
			id = m.tempMsg.ID
		case GroupMessage:
			id = m.groupMsg.ID
			time = m.groupMsg.Time
		}

		content = append([]message.IMessageElement{&message.ReplyElement{
			ReplySeq:  id,
			SenderUin: m.Sender.Uin,
			SenderUID: m.Sender.UID,
			GroupUin:  groupUin,
			Time:      time,
			Elements:  m.Elements,
		}}, content...)
	}

	switch m.Type {
	case PrivateMessage:
		msg, err := m.QQClient.SendPrivateMessage(m.Sender.Uin, content)
		return fromPrivateMessage(m.QQClient, msg), err
	case TempMessage:
		msg, err := m.QQClient.SendTempMessage(m.tempMsg.GroupUin, m.Sender.Uin, content)
		return fromTempMessage(m.QQClient, msg), err
	case GroupMessage:
		msg, err := m.QQClient.SendGroupMessage(m.groupMsg.GroupUin, content)
		return fromGroupMessage(m.QQClient, msg), err
	}
	return nil, fmt.Errorf("invalid type")
}

func (m *MessageContext) ReplyText(text string, quote ...bool) (*MessageContext, error) {
	return m.ReplyMessage([]message.IMessageElement{message.NewText(text)}, quote...)
}

func (m *MessageContext) ReplyImage(img []byte, quote ...bool) (*MessageContext, error) {
	return m.ReplyMessage([]message.IMessageElement{message.NewImage(img)}, quote...)
}

func (m *MessageContext) ReplyFileImage(path string, quote ...bool) (*MessageContext, error) {
	img, err := message.NewFileImage(path)
	if err != nil {
		return nil, err
	}
	return m.ReplyMessage([]message.IMessageElement{img}, quote...)
}

func (m *MessageContext) ReplyRecord(data []byte) (*MessageContext, error) {
	return m.ReplyMessage([]message.IMessageElement{message.NewRecord(data)})
}

func (m *MessageContext) ReplyFileRecord(path string) (*MessageContext, error) {
	record, err := message.NewFileRecord(path)
	if err != nil {
		return nil, err
	}
	return m.ReplyMessage([]message.IMessageElement{record})
}
