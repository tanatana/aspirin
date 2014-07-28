package main

import (
	"fmt"
	"../../aspirin"
)

func main() {
	asp := aspirin.NewAspirin()

	// tmux で例えると `C-t %` => `C-t "`
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.VirticalSplit)
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)
	fmt.Println(*asp)
	for _, window := range asp.GetWindows() {
		fmt.Printf("\t%v\n", *window)
		p := window.GetRootPane()
		aspirin.DrawPaneTree(p)

	}


}
