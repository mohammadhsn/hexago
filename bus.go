package hexago

import "reflect"

// CmdBus decouples Cmd from its execution, dependencies, side effects.
type CmdBus interface {
	Handle(Cmd) string
}

type InMemoryBus struct {
	cmd            *CmdBag
	events         *EventBag
	capturedEvents []*Event
}

func (i InMemoryBus) Handle(cmd Cmd) string {
	hf, err := i.cmd.FactoryFor(cmd.Name)

	if err != nil {
		return ""
	}

	h := hf()

	res, err := h.Handle(cmd)

	v, err := extractRepoField(reflect.ValueOf(h))

	if err != nil {
		return res
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
		i.HandleEvent(*e)
	}

	return res
}

func (i InMemoryBus) HandleEvent(e Event) {
	for _, ef := range i.events.FactoryFor(e.Name) {
		ef().Handle(e)
	}
}
