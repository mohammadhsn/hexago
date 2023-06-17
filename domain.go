package hexago

import (
	"github.com/google/uuid"
	"time"
)

type AggRoot interface {
	Id() uuid.UUID
	Pull() []Event
}

type Event struct {
	Identifier uuid.UUID
	Name       string
	At         time.Time
	Payload    interface{}
}
