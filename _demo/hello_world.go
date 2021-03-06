package main

import (
	"../../aspirin"
	"fmt"
)

func main() {
	asp := aspirin.NewAspirin()
	asp.Debug = true

	w := aspirin.NewWindow("", asp.Width(), asp.Height())
	p := newHelloPane()
	w.SetInitialPane(p)
	w.OnKey(func (ev aspirin.Event){
		// LATIN CAPITAL LETTER 'S'
		if ev.Ch == 83 {
			newPane := newHelloPane()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), newPane, aspirin.PRVirticalSplit)
			asp.DebugPrint("user press 'S'")
		}
		// LATIN SMALL LETTER 's'
		if ev.Ch == 115 {
			newPane := newHelloPane()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), newPane, aspirin.PRHorizontalSplit)
			asp.DebugPrint("user press 's'")
		}
		// LATIN SMALL LETTER 'x'
		if ev.Ch == 120 {
			asp.ActiveWindow().ClosePane(asp.ActiveWindow().ActivePane(), true)
			asp.DebugPrint("user press 'x'")
		}
		// LATIN SMALL LETTER '['
		if ev.Ch == 91 {
			asp.ActiveWindow().MoveToPrevPane()
			asp.DebugPrint("user press '['")
		}
		// LATIN SMALL LETTER ']'
		if ev.Ch == 93 {
			asp.ActiveWindow().MoveToNextPane()
			asp.DebugPrint("user press ']'")
		}
	})
	// w.Init()

	asp.AddWindow(w, true)

	asp.OnKey(func(ev aspirin.Event){
		// LATIN SMALL LETTER 'q'
		if ev.Ch == 113 {
			asp.Quit()
		}
		// LATIN SMALL LETTER '{'
		if ev.Ch == 123 {
			asp.MoveToPrevWindow()
		}
		// LATIN SMALL LETTER '}'
		if ev.Ch == 125 {
			asp.MoveToNextWindow()
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
		if 48 <= ev.Ch && ev.Ch <= 57 {
			// line := aspirin.NewTextLine(fmt.Sprintf("%v (active: %v)", ev, p))
			// line := aspirin.NewTextLine(fmt.Sprintf("%v::%v", p.Id(), p.Size()))
			line := aspirin.NewTextLine("Hello, world!")
			p.AddLine(line, true)
		}
		if ev.Ch == 106 {
			p.MoveNextLine()
		}
		if ev.Ch == 107 {
			p.MovePrevLine()
		}
	})

	p.OnResize(func(ev aspirin.Event){
		line := aspirin.NewTextLine(fmt.Sprintf("terminal resized (%v)", ev))
		p.AddLine(line, false)
	})

	return p
}

func (hp *HelloPane)ViewDidLoad() {
	hp.AddLine(aspirin.NewTextLine("Hello, world! DEMO"), false)
	hp.AddLine(aspirin.NewTextLine(""), false)

	hp.AddLine(aspirin.NewTextLine("insert something: press '0-9' key)"), false)
	hp.AddLine(aspirin.NewTextLine("virtical split  : press 'S' key (Shift-s)"), false)
	hp.AddLine(aspirin.NewTextLine("horizontal split: press 's' key"), false)
	hp.AddLine(aspirin.NewTextLine("scroll up  : press 'j' key (Shift-s)"), false)
	hp.AddLine(aspirin.NewTextLine("scroll down: press 'k' key"), false)

	hp.AddLine(aspirin.NewTextLine("when there are some pains,"), false)
	hp.AddLine(aspirin.NewTextLine("  move to next pain  : press ']' key (Shift-s)"), false)
	hp.AddLine(aspirin.NewTextLine("  move to prev pain  : press '[' key"), false)


}
