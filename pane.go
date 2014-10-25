package aspirin

import (
	"fmt"
	"github.com/nsf/termbox-go"
)


type Pane interface {
	viewDidLoad()
	onKey()
	getEventChannel() chan Event
	setupEventLoop()
}

type BasePane struct{
	id int
	parent Pane
	left, right Pane
	EventChannel chan Event
}

func (bp *BasePane)viewDidLoad() {
	fmt.Printf("viewDidLoad from %s\n", "BasePane")
}

func (bp *BasePane)onKey(ev Event) {
	fmt.Printf("onKey from %s\n", "BasePane")
}

func (bp *BasePane)getEventChannel() chan Event{
	return bp.EventChannel
}

func (bp *BasePane)setupEventLoop() {
	loop:

	for {
		ev := <- bp.EventChannel
		fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go bp.onKey(ev)
		case EventQuit:
			fmt.Printf("EventQuit was handled\n");
			break loop
		}
	}
}
