package hexago

import "github.com/google/uuid"

type AggRoot interface {
	Id() uuid.UUID
}
