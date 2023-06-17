package hexago

import (
	"github.com/google/uuid"
	"time"
)

// AggRoot is just an abbreviation for AggregateRoot. In DDD terms it clusters domain
// entities so that they act like one single unit. It has ability to capturing events.
// The brain of the business and handling core considerations goes here in AggRoot.
type AggRoot interface {
	Identifiable
	Pull() []Event
}

// Event is like an important happening in the domain. It can be implied a change state,
// or something like that.
type Event struct {
	Identifier uuid.UUID
	Name       string
	At         time.Time
	Payload    interface{}
}

// Repo is responsible take and bring the domain objects and AggRoot's.
type Repo interface {
	Seen() []AggRoot
}
