package event

import "time"

type Message interface {
	Key() string
}

const (
	CREATED = "meow.created"
)

type MeowCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

func (m *MeowCreatedMessage) Key() string {
	return CREATED
}
