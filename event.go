package hexago

import "github.com/google/uuid"

type EventHandler interface {
	Handle(Event)
}

// EventBag just holds a mapping between Event and EventHandler
type EventBag struct {
	events map[string][]EventHandlerFactory
}

// NewEventBag creates an empty EventBag.
func NewEventBag() *EventBag {
	return &EventBag{make(map[string][]EventHandlerFactory)}
}

func (e *EventBag) Add(event string, handlers ...EventHandlerFactory) {
	e.events[event] = append(e.events[event], handlers...)
}

func (e *EventBag) FactoryFor(event string) []EventHandlerFactory {
	return e.events[event]
}

type EventHandlerFactory func() EventHandler

func (e Event) Id() uuid.UUID {
	return e.Identifier
}
