package main

import (
	"../../aspirin"
	"fmt"
)

func main() {
	asp := aspirin.NewAspirin()

	w := aspirin.NewWindow("")
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

}
