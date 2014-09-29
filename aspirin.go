package aspirin

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type Aspirin struct {
	activeWindow int
	windows []*window
	windowCounter int
	width, height int
}

func NewAspirin(width int, height int) *Aspirin {
	ap := new(Aspirin)
	ap.windowCounter = 0
	ap.width         = width
	ap.height        = height
	ap.CreateWindow("window")

	return ap
}

func (ap *Aspirin)CreateWindow(title string) *window{
	w := newWindow(ap.windowCounter, title, ap.width, ap.height)
	ap.windows = append(ap.windows, w)
	ap.activeWindow = w.id
	ap.windowCounter += 1
	return w
}

func (ap *Aspirin)GetWindows() []*window{
	return ap.windows
}

func (ap *Aspirin)GetActiveWindow() *window{
	return ap.windows[ap.activeWindow]
}

func (ap *Aspirin)Draw() *window{
	return ap.windows[ap.activeWindow]
}

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

// print aspirin state for debugging
func DrawPaneTree(targetPane *pane, floor int) {
	fmt.Print("  ")
	for i := 0; i < floor; i++ {
		fmt.Print("  ")
	}
	printf_tb(0, targetPane.id + 1, termbox.ColorWhite, termbox.ColorBlack, "%v", *targetPane)
	termbox.Flush()

	if (targetPane.left != nil){
		DrawPaneTree(targetPane.left, floor + 1)
	}
	if (targetPane.right != nil){
		DrawPaneTree(targetPane.right, floor + 1)
	}

}
