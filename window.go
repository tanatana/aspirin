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

func NewWindow(title string, width, height int) *window{
	w             := new(window)
	w.title        = title
	w.width        = width
	w.height       = height

	w.eventChannel = make(chan Event)
	w.onKey = (func(e Event){})

	p := new(RootPane)
	p.Init()
	p.SetSize(0, 0, width, height)
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

func (w *window)Width() int{
	return w.width
}

func (w *window)Height() int{
	return w.height
}

func (w *window)MoveToNextPane() {}
func (w *window)MoveToPrevPane() {}
func (w *window)MoveToPane() {}

func (w *window)SetInitialPane(child Pane, changeActivePane bool) {
	child.setParent(w.rootPane)
	w.rootPane.setLeft(child)
	if (changeActivePane) {
		w.activePane = child
	}
}
