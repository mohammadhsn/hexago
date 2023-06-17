package hexago

import "github.com/google/uuid"

type Identifiable interface {
	Id() uuid.UUID
}

func NewId() uuid.UUID {
	return uuid.New()
}
