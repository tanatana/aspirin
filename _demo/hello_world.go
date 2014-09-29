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
	// asp.CreateWindow("test")
	// tmux で例えると `C-t %` => `C-t "` => `C-t o` => `C-t "`
	// asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.VirticalSplit)
	//
	// asp.GetActiveWindow().SetActivePane(1)
	// asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)
	asp.Reflesh()
	asp.DrawStatus()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			printf_tb(0, 20, termbox.ColorDefault, termbox.ColorDefault, "%s", string(ev.Ch))
			printf_tb(0, 21, termbox.ColorDefault, termbox.ColorDefault, "%d", ev.Ch)
			if ev.Ch == 34 {
				// HorizontalSplit
				asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.HorizontalSplit)
			}
			if ev.Ch == 37 {
				// VerticalSplit
				asp.GetActiveWindow().SplitPane(asp.GetActiveWindow().GetActivePane(), aspirin.VirticalSplit)
			}
			if ev.Ch == 100 {
				// close active pane
			}
			if ev.Ch == 113 {
				break loop
			}
			asp.Reflesh()
			asp.DrawStatus()
		}
	}
}
