package event

import "time"

type Event struct {
	AggregateID   string
	AggregateType string
	Type          string
	Data          []byte
	MetaData      []byte
	Version       uint32
	Timestamp     time.Time
}
