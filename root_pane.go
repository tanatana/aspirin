package aspirin

import (
	"fmt"
)

type RootPane struct {
	BasePane
}

func (rp *RootPane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "RootPane")
}

func (rp *RootPane)onKey(ev Event) {
	fmt.Printf("onKey@%s\n", "RootPane")
}

func (rp *RootPane)SetRight(p Pane) {
	panic("Root Pane can not have right pane")
}
