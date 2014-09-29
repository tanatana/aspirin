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
	debug bool
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
func (ap *Aspirin)DrawStatus() {
	drawLine := 0
	var drawPaneTree func(targetPane *pane, floor int)
	drawPaneTree = func (targetPane *pane, floor int) {
		printf_tb((floor + 1) * 2, drawLine, termbox.ColorWhite, termbox.ColorBlack, "%v", *targetPane)
		drawLine += 1
		if (targetPane.left != nil){
			drawPaneTree(targetPane.left, floor + 1)
		}
		if (targetPane.right != nil){
			drawPaneTree(targetPane.right, floor + 1)
		}
	}

	for _, window := range ap.GetWindows() {
		printf_tb(0, drawLine, termbox.ColorWhite, termbox.ColorBlack, "%v", *window)
		drawLine += 1
		p := window.GetRootPane()
		drawPaneTree(p, 1)
	}

	termbox.Flush()
}
