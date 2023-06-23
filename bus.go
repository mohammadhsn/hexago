package hexago

import "reflect"

// CmdBus decouples Cmd from its execution, dependencies, side effects.
type CmdBus interface {
	Handle(Cmd) (string, error)
}

type InMemoryBus struct {
	cmd            *CmdBag
	events         *EventBag
	capturedEvents []*Event
}

// NewInMemoryBus creates a fresh bus without any captured events.
func NewInMemoryBus(cmd *CmdBag, events *EventBag) InMemoryBus {
	return InMemoryBus{
		cmd:            cmd,
		events:         events,
		capturedEvents: make([]*Event, 0),
	}
}

func (i InMemoryBus) Handle(cmd Cmd) (string, error) {
	hf, err := i.cmd.FactoryFor(cmd.Name)

	if err != nil {
		return "", err
	}

	h := hf()

	res, err := h.Handle(cmd)

	v, err := extractRepoField(reflect.ValueOf(h))

	if err != nil {
		return res, err
	}

	repo, ok := v.Interface().(Repo)

	if ok {
		for _, agg := range repo.Seen() {
			for _, e := range agg.Pull() {
				i.capturedEvents = append(i.capturedEvents, &e)
			}
		}
	}

	for len(i.capturedEvents) != 0 {
		e := i.capturedEvents[0]
		i.capturedEvents = i.capturedEvents[1:]
		i.handleEvent(*e)
	}

	return res, nil
}

func (i InMemoryBus) handleEvent(e Event) {
	for _, ef := range i.events.FactoryFor(e.Name) {
		ef().Handle(e)
	}
}
