package main

import (
	"../../aspirin"
	"fmt"
)

func main() {
	asp := aspirin.NewAspirin()

	w := aspirin.NewWindow("", asp.Width(), asp.Height())
	p := new(MainPane)
	p.Init()
	p.OnKey(func(ev aspirin.Event){
		// aspirin.Printf_tb(0, 0, termbox.ColorDefault, termbox.ColorDefault, "onKey@%s\n", "MainPane")
		// termbox.Flush()
		if ev.Ch == 0 {
			p.ActiveLine().RunAction(ev)
		} else if ev.Ch == 106 {
			p.MoveNextElement()
		} else if ev.Ch == 107 {
			p.MovePrevElement()
		} else {
			lo := new(aspirin.LineBase)
			lo.SetText(fmt.Sprintf("%v", lo))

			lo.SetAction(func(e aspirin.Event){
				loFromAction := aspirin.NewTextLine(fmt.Sprintf("%v from action", ev))
				p.AddLine(loFromAction, false)
			})

			p.AddLine(lo, false)
		}

	})


	p.SetSize(0, 0, w.Width(), w.Height())
	w.SetInitialPane(p, true)
	asp.AddWindow(w, true)

	asp.OnKey(func(ev aspirin.Event){
		if ev.Ch == 113 {
			asp.Quit()
		}
	})

	asp.Run()
	fmt.Printf("%v\n", asp)
	fmt.Printf("%v\n", asp.ActiveWindow())
	fmt.Printf("%v\n", asp.ActiveWindow().RootPane())
	fmt.Printf("%v\n", asp.ActiveWindow().RootPane().Left())
}

type MainPane struct {
	aspirin.BasePane
}
