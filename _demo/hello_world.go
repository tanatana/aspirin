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
	p.OnKey(func(ev aspirin.Event) {
		line := aspirin.NewTextLine("Hello, world")
		p.AddLine(line, false)
	})
	p.SetSize(0, 0, w.Width(), w.Height())
	w.SetInitialPane(p, true)
	asp.AddWindow(w, true)

	asp.OnKey(func(ev aspirin.Event){
		if ev.Ch == 113 {
			asp.Quit()
		}
		if ev.Ch == 115 {
			asp.DebugPrint("split")

			tmpPane := new(aspirin.BasePane)
			tmpPane.Init()
			asp.ActiveWindow().SplitPane(asp.ActiveWindow().ActivePane(), tmpPane, 0)
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
