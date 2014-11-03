package main

import (
	"../../aspirin"
	"fmt"
)

func main() {
	asp := aspirin.NewAspirin()

	w := aspirin.NewWindow("", asp.Width(), asp.Height())
	p := newHelloPane()
	w.SetInitialPane(p, true)
	asp.AddWindow(w, true)

	asp.OnKey(func(ev aspirin.Event){
		if ev.Ch == 113 {
			asp.Quit()
		}


		if ev.Ch == 83 {
			newPane := newHelloPane()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), newPane, aspirin.SplitVirtical)
		}
		if ev.Ch == 115 {
			newPane := newHelloPane()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), newPane, aspirin.SplitHorizontal)
		}
	})

	asp.Run()
}

type HelloPane struct {
	aspirin.BasePane
}

func newHelloPane() aspirin.Pane{
	p := new(HelloPane)
	p.Init()
	p.OnKey(func(ev aspirin.Event) {
		line := aspirin.NewTextLine(fmt.Sprintf("Hello, world (%v)", ev))
		p.AddLine(line, false)
	})
	return p
}
