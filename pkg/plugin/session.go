package plugin

import "github.com/google/uuid"

type Session struct {
	app *App
	ids []string
}

func (s *Session) Close() {
	app := s.app
	for _, id := range s.ids {
		if _, ok := app.messageHandlers[id]; ok {
			delete(app.messageHandlers, id)
		}
		if _, ok := app.commandHandlers[id]; ok {
			delete(app.commandHandlers, id)
		}
	}
}

func (s *Session) newID() string {
	id := uuid.New().String()
	s.ids = append(s.ids, id)
	return id
}

func (s *Session) OnMessage(pattern string) *EventHandle[*MessageContext] {
	id := s.newID()
	s.app.messageHandlers[id] = buildMessageHandle(pattern)
	return s.app.messageHandlers[id]
}

func (s *Session) OnCommand(prefix string) *EventHandle[*CommandContext] {
	id := s.newID()
	s.app.commandHandlers[id] = buildCommandHandle(prefix)
	return s.app.commandHandlers[id]
}
