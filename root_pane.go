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
	fmt.Printf("viewDidLoad from %s\n", "RootPane")
}

func (rp *RootPane)onKey() {
	fmt.Printf("onKey from %s\n", "RootPane")
}

func (rp *RootPane)getEventChannel() chan Event{
	return rp.EventChannel
}
