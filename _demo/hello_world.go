package main

import (
	"fmt"
	"../../aspirin"
)

func main() {
	asp := aspirin.NewAspirin()
	// asp.CreateWindow()
 	// asp.ActiveWindow().CreatePane(0)
 	// asp.ActiveWindow().CreatePane(0)
 	// asp.ActiveWindow().CreatePane(0)
	fmt.Println(asp)
	for _, window := range asp.GetWindows() {
		fmt.Printf("\t%v\n", window)
		p :=window.GetRootPane()
		fmt.Printf("\t\t%v\n", p)
	}

}
