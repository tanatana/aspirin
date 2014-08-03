package main

import (
	"fmt"
	"../../aspirin"
)

func main() {
	asp := aspirin.NewAspirin()

	// tmux で例えると `C-t %` => `C-t "` => `C-t o` => `C-t "`
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.VirticalSplit)
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)
	asp.GetActiveWindow().SetActivePane(1)
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)


	for _, window := range asp.GetWindows() {
		fmt.Printf("%v\n", *window)
		p := window.GetRootPane()
		aspirin.DrawPaneTree(p, 0)
	}

}
