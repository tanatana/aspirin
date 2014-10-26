package aspirin

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type aspirin struct {
	activeWindow *window
	windows []*window
	windowCounter int
	width, height int
	onKey func(ev Event)
	EventChannel chan Event
}

func (asp *aspirin)OnKey(f func(ev Event)){
	asp.onKey = f
}

func (asp *aspirin)AddWindow(w *window, changeActiveWindow bool) {
	w.id = asp.windowCounter
	if w.title == "" {
		w.title = fmt.Sprintf("window %v", w.id)
	}

	asp.windows = append(asp.windows, w)
	if (changeActiveWindow) {
		asp.activeWindow = w
	}
	asp.windowCounter += 1
}

func (asp *aspirin)ActiveWindow() *window{
	return asp.activeWindow
}


func (asp *aspirin)Run(){
	go setupEventLoop(asp.EventChannel)

loop:
	for {
		ev := <- asp.EventChannel
		fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go asp.onKey(ev)
		case EventQuit:
			fmt.Printf("EventQuit was handled\n");
			break loop
		}
		asp.activeWindow.eventChannel <- ev
	}
}

func (asp *aspirin)Quit(){
	termbox.Close()
	var e Event
	e.Type = EventQuit
	asp.EventChannel <- e
}

func setupEventLoop(ec chan Event) {
	for {
		ev := termbox.PollEvent()
		ec <- NewEventWithTermboxEvent(ev)
	}
}

func (asp *aspirin)Width() int{
	return asp.width
}

func (asp *aspirin)Height() int{
	return asp.height
}


func NewAspirin() *aspirin{
	asp := new(aspirin)
	asp.EventChannel = make(chan Event)
	asp.onKey = (func(ev Event){})

	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}

	asp.width, asp.height = termbox.Size()

	return asp
}


func NewAspirinApp() *aspirin{
	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	defer termbox.Close()
	width, height := termbox.Size()

	ap := new(aspirin)
	ap.width         = width
	ap.height        = height
	ap.windowCounter = 0
	// ap.CreateWindow("window")

	return ap
}
