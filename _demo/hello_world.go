package main

import (
	"../../aspirin"
	"fmt"
)

func main() {
	asp := aspirin.NewAspirin()

	w := aspirin.NewWindow("", asp.Width(), asp.Height())
	p := newHelloPane()
	w.SetInitialPane(p)
	asp.AddWindow(w, true)

	asp.OnKey(func(ev aspirin.Event){
		// LATIN SMALL LETTER 'q'
		if ev.Ch == 113 {
			asp.Quit()
		}
		// LATIN CAPITAL LETTER 'S'
		if ev.Ch == 83 {
			newPane := newHelloPane()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), newPane, aspirin.PRVirticalSplit)
		}
		// LATIN SMALL LETTER 's'
		if ev.Ch == 115 {
			newPane := newHelloPane()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), newPane, aspirin.PRHorizontalSplit)
		}
		// LATIN SMALL LETTER '['
		if ev.Ch == 91 {
			asp.ActiveWindow().MoveToPrevPane()
		}
		// LATIN SMALL LETTER ']'
		if ev.Ch == 93 {
			asp.ActiveWindow().MoveToNextPane()
		}
		// LATIN SMALL LETTER 'x'
		if ev.Ch == 120 {
			// asp.ActiveWindow().ClosePane(asp.ActiveWindow().ActivePane())
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

	p.OnResize(func(ev aspirin.Event){
		line := aspirin.NewTextLine(fmt.Sprintf("terminal resized (%v)", ev))
		p.AddLine(line, false)
	})

	return p
}


// 名前空間違うから動かないよこれ
func (hp *HelloPane)viewDidLoad() {
	line := aspirin.NewTextLine(fmt.Sprintf("Hello, World! (%v)", hp))
	hp.AddLine(line, false)
}
