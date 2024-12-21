package plugin

type EventHandle[T any] struct {
	matcher func(ctx T) bool
	handler func(ctx T)
}

func (h *EventHandle[T]) Handle(handler func(ctx T)) *EventHandle[T] {
	h.handler = handler
	return h
}

func (h *EventHandle[T]) match(e T) bool {
	if h.matcher != nil {
		return h.matcher(e)
	}
	return true
}

func (h *EventHandle[T]) dispatch(ctx T) {
	if h.handler != nil {
		h.handler(ctx)
	}
}
