package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"../../aspirin"
)

func main() {
	// termbox の初期化
	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	defer termbox.Close()
	width, height := termbox.Size()

	asp := aspirin.NewAspirin(width, height)
	// tmux で例えると `C-t %` => `C-t "` => `C-t o` => `C-t "`
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.VirticalSplit)
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)
	asp.GetActiveWindow().SetActivePane(1)
	asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)

	fmt.Printf("%s\n", "press 'q' to quit")
	for _, window := range asp.GetWindows() {
		fmt.Printf("%v\n", *window)
		p := window.GetRootPane()
		aspirin.DrawPaneTree(p, 0)
	}

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 113 {
				break loop
			}
		}
	}
}
