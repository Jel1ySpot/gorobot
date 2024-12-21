package plugin

import (
	"fmt"
	gorobot "github.com/Jel1ySpot/gorobot/pkg/bot"
	"github.com/google/uuid"
)

type App struct {
	name   string
	Logger gorobot.Logger

	loadEvent       *EventHandle[*gorobot.Bot]
	messageHandlers map[string]*EventHandle[*MessageContext]
	commandHandlers map[string]*EventHandle[*CommandContext]
}

func Create(name string) *App {
	app := &App{
		name:   name,
		Logger: gorobot.NewLogger(fmt.Sprintf("%s -> ", name)),

		loadEvent:       &EventHandle[*gorobot.Bot]{},
		messageHandlers: make(map[string]*EventHandle[*MessageContext]),
		commandHandlers: make(map[string]*EventHandle[*CommandContext]),
	}
	apps = append(apps, app)
	app.Logger.Info("Plugin created")
	return app
}

func (a *App) OnLoad(callback func(bot *gorobot.Bot)) *EventHandle[*gorobot.Bot] {
	return a.loadEvent.Handle(callback)
}

func (a *App) OnMessage(pattern string) *EventHandle[*MessageContext] {
	id := uuid.New().String()
	a.messageHandlers[id] = buildMessageHandle(pattern)
	return a.messageHandlers[id]
}

func (a *App) OnCommand(prefix string) *EventHandle[*CommandContext] {
	id := uuid.New().String()
	a.commandHandlers[id] = buildCommandHandle(prefix)
	return a.commandHandlers[id]
}

func (a *App) CreateSession() *Session {
	session := &Session{
		app: a,
	}
	return session
}
