package aspirin

import (
	// "fmt"
)

type debugPane struct {
	BasePane
}

func newDebugPane() Pane{
	p := new(debugPane)
	p.Init()

	return p
}
