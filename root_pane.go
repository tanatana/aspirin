package aspirin

import (
	"fmt"
)

type RootPane struct {
	BasePane
}

func newRootPane() Pane{
	rp := new(RootPane)
	rp.EventChannel = make(chan Event)
	go rp.setupEventLoop()

	return rp
}

func (rp *RootPane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "RootPane")
}

func (rp *RootPane)onKey() {
	fmt.Printf("onKey@%s\n", "RootPane")
}

func (rp *RootPane)getEventChannel() chan Event{
	return rp.EventChannel
}
