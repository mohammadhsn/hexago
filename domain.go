package hexago

import (
	"github.com/google/uuid"
	"time"
)

type AggRoot interface {
	Identifiable
	Pull() []Event
}

type Event struct {
	Identifier uuid.UUID
	Name       string
	At         time.Time
	Payload    interface{}
}

type Repo interface {
	Seen() []AggRoot
}
