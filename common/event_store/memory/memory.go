package memory

import (
	"errors"
	"github.com/Igor-Rashyn/project-query-service/common/event"
	"sync"
)

var ErrEventVersion = errors.New("event version is old")

type Memory struct {
	events          map[string][]event.Event
	currentVersions map[string]uint32
	lock            sync.Mutex
}

func (e *Memory) Save(event event.Event) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	streamName := streamKey(event.AggregateType, event.AggregateID)

	stream := e.events[streamName]

	var currentVersion uint32

	if len(stream) > 0 {
		lastEvent := stream[len(stream)-1]
		currentVersion = lastEvent.Version
	}

	if event.Version != currentVersion+1 {
		return ErrEventVersion
	}

	stream = append(stream, event)

	e.events[streamName] = stream

	return nil
}

//streamKey generates an stream key to store events against from aggregateType and aggregateID
func streamKey(aggregateType, aggregateID string) string {
	return aggregateType + "_" + aggregateID
}
