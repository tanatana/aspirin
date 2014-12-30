package aspirin

import (
	"github.com/nsf/termbox-go"
)

// Event type. See Event.Type field.
const (
	EventQuit termbox.EventType = 100
)

type Event struct{
	termbox.Event
}

func NewEvent() *Event{
	return new(Event)
}

func NewEventWithTermboxEvent(ev termbox.Event) Event{
	var e Event
	e.Event = ev
	return e
}
