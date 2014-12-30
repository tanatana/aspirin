package aspirin

import (
	// "fmt"
)

type debugPane struct {
	BasePane
}

func (dp *debugPane)AddLine(lo Line, setActive bool) {
	dp.lines = append(dp.lines, lo)

	if setActive {
		//bp.activeLine = lo
	}
}

func newDebugPane() Pane{
	p := new(debugPane)
	p.Init()

	return p
}
