package aspirin

import (
	"fmt"
)

type RootPane struct {
	BasePane
}

func newRootPane(id int, width, height int) Pane{
	p := new(RootPane)
	p.Init()
	p.setId(id)
	p.setSize(0, 0, width, height)
	p.setRole(PRRoot)
	return p
}

func (rp *RootPane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "RootPane")
}

func (rp *RootPane)SetRight(p Pane) {
	panic("Root Pane can not have right pane")
}
