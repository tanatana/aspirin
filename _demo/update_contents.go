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
		tlo := new(aspirin.TextLineObject)
		tlo.SetText(fmt.Sprintf("%v", ev))
		p.AddLineObject(tlo)
	})


	// p.contents := [] aspirin.LineObject
	// p.SetContents(contents)

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
