package aspirin

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Pane interface {
	viewDidLoad()
	onKey()
	EventChannel() chan Event
	setupEventLoop()
	Split(SplitType)
	Close()
}

type BasePane struct{
	id int
	parent Pane
	left, right Pane
	x, y int
	width, height int
	eventChannel chan Event
}

type SplitType int

const (
	VirticalSplit SplitType = iota
	HorizontalSplit
)

func (bp *BasePane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "BasePane")
}

func (bp *BasePane)onKey(ev Event) {
	fmt.Printf("onKey@%s\n", "BasePane")
}

func (bp *BasePane)Split(t SplitType) {}

func (bp *BasePane)Close() {}

func (bp *BasePane)EventChannel() chan Event{
	return bp.eventChannel
}

func (bp *BasePane)setupEventLoop() {
	for {
		ev := <- bp.eventChannel
		fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go bp.onKey(ev)
		}
	}
}
