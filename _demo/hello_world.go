package main

import (
	"../../aspirin"
	"fmt"
)

func main() {
	asp := aspirin.NewAspirin()

	asp.OnKey(func(ev aspirin.Event){
		if ev.Ch == 113 {
			asp.Quit()
		}
	})

	asp.Run()

	fmt.Printf("%v\n", asp)
}
