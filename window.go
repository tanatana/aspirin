package aspirin

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type window struct {
	id int
	activePane Pane
	title string
	width, height int
	rootPane Pane
	onKey func(ev Event)
	eventChannel chan Event
}

func (w *window)setupEventLoop() {
	for {
		ev := <- w.eventChannel
		fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go w.onKey(ev)
		}
		w.activePane.EventChannel() <- ev
	}
}

func NewWindow(title string) *window{
	// TODO: create mini-pane like emacs's mini-buffer
	w             := new(window)
	w.title        = title
	w.eventChannel = make(chan Event)

	w.onKey = (func(e Event){})

	p := newRootPane()
	w.rootPane = p
	w.activePane = p

	go w.setupEventLoop()

	return w
}

func (w *window)RootPane() Pane{

	return w.rootPane
}

func (w *window)ActivePane() Pane{

	return w.activePane
}
