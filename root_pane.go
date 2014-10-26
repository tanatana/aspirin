package aspirin

import (
	"fmt"
)

type RootPane struct {
	BasePane
}

func newRootPane(x, y, width, height int) Pane{
	rp := new(RootPane)
	rp.eventChannel = make(chan Event)
	go rp.setupEventLoop()

	rp.x = x
	rp.y = y
	rp.width = width
	rp.height = height

	return rp
}

func (rp *RootPane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "RootPane")
}

func (rp *RootPane)onKey() {
	fmt.Printf("onKey@%s\n", "RootPane")
}
