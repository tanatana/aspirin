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
	eventChannel chan termbox.Event
}

type Event struct{
	termbox.Event
}

func initTermbox(ch chan termbox.Event, done chan bool) {
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
			ch <- ev
			if ev.Ch == 113 {
				fmt.Printf("%v\n", ev)
				break loop
			}
		}
	}
	termbox.Close()
	close(ch)
	done <- true
}

func NewAspirin() *aspirin{
	asp := new(aspirin)
	asp.eventChannel = make(chan termbox.Event)
	done := make(chan bool)
	go initTermbox(asp.eventChannel, done)
	for {
		ev := <- asp.eventChannel
		fmt.Printf("%v\n", ev)
		if ev.Ch == 113 {
			break
		}

	}
	<- done
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
