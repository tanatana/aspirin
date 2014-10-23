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
	eventChannel chan Event
	done chan bool
	onKey func(ev Event)
}

func (asp *aspirin)OnKey(f func(ev Event)){
	asp.onKey = f
}

func (asp *aspirin)Quit(){
	termbox.Close()
	asp.done <- true
}

func setupTermbox(ch chan Event, done chan bool) {
	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	defer termbox.Close()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			ch <- NewEventWithTermboxEvent(ev)
			if ev.Ch == 113 {
				break loop
			}
		}
	}


	termbox.Close()
	done <- true
}

func NewAspirin() *aspirin{
	asp := new(aspirin)
	asp.eventChannel = make(chan Event)
	asp.done = make(chan bool)
	go setupTermbox(asp.eventChannel, asp.done)

	for {
		ev := <- asp.eventChannel
		fmt.Printf("%v\n", ev)
		fmt.Printf("%v\n", asp.onKey)
		if ev.Ch == 113 {
			break
		}
	}

	<- asp.done
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
