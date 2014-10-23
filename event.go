package aspirin

import (
	"github.com/nsf/termbox-go"
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
