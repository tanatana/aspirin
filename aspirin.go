package aspirin

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type aspirin struct {
	activeWindowId int
	windows []*Window
	windowCounter int
	width, height int
	onKey func(ev Event)
	EventChannel chan Event
}

func (asp *aspirin)OnKey(f func(ev Event)){
	asp.onKey = f
}

func (asp *aspirin)Run(){
	fmt.Printf("asp.Run() was called()");
	go setupEventLoop(asp.EventChannel)

	fmt.Printf("2nd phase in asp.Run()");

loop:
	for {
		ev := <- asp.EventChannel
		fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go asp.onKey(ev)
		case EventQuit:
			fmt.Printf("EventQuit was handled");
			break loop
		}
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

func NewAspirin() *aspirin{
	asp := new(aspirin)
	asp.EventChannel = make(chan Event)

	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}

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
