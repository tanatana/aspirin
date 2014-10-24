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
	quit chan bool
	onKey func(ev Event)
}

func (asp *aspirin)OnKey(f func(ev Event)){
	asp.onKey = f
}

func (asp *aspirin)Run(){
	fmt.Printf("asp.Run() was called()");
	go setupEventLoop(asp.eventChannel, asp.quit)

	fmt.Printf("2nd phase in asp.Run()");
	for {
		ev := <- asp.eventChannel
		switch ev.Type {
		case termbox.EventKey:
			fmt.Printf("%v\n",ev)
			asp.onKey(ev)
		}
	}
	fmt.Printf("3rd phase in asp.Run()");
	<- asp.quit
}

func (asp *aspirin)Quit(){
	termbox.Close()
	asp.quit <- true
}

func setupEventLoop(ch chan Event, quit chan bool) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			ch <- NewEventWithTermboxEvent(ev)
		}
	}
}

func NewAspirin() *aspirin{
	asp := new(aspirin)
	asp.eventChannel = make(chan Event)
	asp.quit = make(chan bool)

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
