package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"../../aspirin"
)

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}



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


	for _, window := range asp.GetWindows() {
		fmt.Printf("%v\n", *window)
		p := window.GetRootPane()
		aspirin.DrawPaneTree(p, 0)
	}

	number := 0
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			number += 1
			printf_tb(0, 13 + number, termbox.ColorWhite, termbox.ColorBlack, "%s", string(ev.Ch))
			printf_tb(1, 13 + number, termbox.ColorWhite, termbox.ColorBlack, ": %d", ev.Ch)
			termbox.Flush()

			if ev.Ch == 113 {
				break loop
			}
		}
	}
}
