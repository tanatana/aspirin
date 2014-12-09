package aspirin

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type aspirin struct {
	activeWindow *window
	windows []*window
	windowCounter int
	width, height int
	onKey func(ev Event)
	onResize func(ev Event)
	onMouse func(ev Event)
	onError func(ev Event)
	beforeQuite func(ev Event)
	EventChannel chan Event


	Debug bool
	debugWindow *window
}

func (asp *aspirin)OnKey(f func(ev Event)){
	asp.onKey = f
}
func (asp *aspirin)OnResize(f func(ev Event)){
	asp.onResize = f
}
func (asp *aspirin)OnMouse(f func(ev Event)){
	asp.onMouse = f
}
func (asp *aspirin)OnError(f func(ev Event)){
	asp.onError = f
}

func (asp *aspirin)Windows() []*window {
	return asp.windows
}
func (asp *aspirin)AddWindow(w *window, changeActiveWindow bool) {
	w.id = asp.windowCounter
	if w.title == "" {
		w.title = fmt.Sprintf("window %v", w.id)
	}

	asp.windows = append(asp.windows, w)
	if (changeActiveWindow) {
		asp.activeWindow = w
	}
	asp.windowCounter += 1
}
func (asp *aspirin)MoveToWindow(target *window){
	asp.activeWindow = target

	asp.RefleshScreen()
}

func (asp *aspirin)MoveToNextWindow() *window{
	nextWin := asp.findNextWindow(asp.activeWindow)
	asp.MoveToWindow(nextWin)
	return nextWin
}
func (asp *aspirin)MoveToPrevWindow() *window{
	prevWin := asp.findPrevWindow(asp.activeWindow)
	asp.MoveToWindow(prevWin)
	return prevWin
}

func (asp *aspirin)findNextWindow(targetWin *window) *window{
	winsMaxIndex := len(asp.windows) - 1

	for index, win := range asp.windows {
		if win == targetWin {
			if index == winsMaxIndex {
				return asp.windows[0]
			}
			return asp.windows[index + 1]
		}
	}
	panic("oops! something wrong!")
}
func (asp *aspirin)findPrevWindow(targetWin *window)  *window{
	winsMaxIndex := len(asp.windows) - 1

	for index, win := range asp.windows {
		if win == targetWin {
			if index == 0 {
				return asp.windows[winsMaxIndex]
			}
			return asp.windows[index - 1]
		}
	}
	panic("oops! something wrong!")
}
func (asp *aspirin)ActiveWindow() *window{
	return asp.activeWindow
}


func (asp *aspirin)Run(){
	if asp.Debug {
		dw := NewWindow("debug", asp.Width(), asp.Height())
		p := newDebugPane()
		dw.SetInitialPane(p)
		dw.Init()
		asp.AddWindow(dw, false)
		asp.debugWindow = dw
	}

	go setupEventLoop(asp.EventChannel)

loop:
	for {
		ev := <- asp.EventChannel
		// fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go asp.onKey(ev)
		case termbox.EventResize:
			if (asp.onResize != nil){
				go asp.onResize(ev)
			}
		case termbox.EventMouse:
			go asp.onMouse(ev)
		case termbox.EventError:
			go asp.onError(ev)
		case EventQuit:
			// fmt.Printf("EventQuit was handled\n");
			if asp.beforeQuite != nil {
				asp.beforeQuite(ev);
			}
			break loop
		}
		asp.activeWindow.eventChannel <- ev
	}
}

func (asp *aspirin)RefleshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	asp.ActiveWindow().displayPaneList.ForEach(func (p Pane) {p.Update()})
}

func (asp *aspirin)Quit(){
	termbox.Close()
	var e Event
	e.Type = EventQuit
	asp.EventChannel <- e
}

func setupEventLoop(ec chan Event) {
	for {
		ev := termbox.PollEvent()
		ec <- NewEventWithTermboxEvent(ev)
	}
}

func (asp *aspirin)Width() int{
	return asp.width
}

func (asp *aspirin)Height() int{
	return asp.height
}

func NewAspirin() *aspirin{
	asp := new(aspirin)
	asp.EventChannel = make(chan Event)
	asp.onKey = (func(ev Event){})

	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}

	asp.width, asp.height = termbox.Size()

	return asp
}

func NewAspirinApp() *aspirin{
	err := termbox.Init()
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	defer termbox.Close()
	width, height := termbox.Size()

	ap := new(aspirin)
	ap.width         = width
	ap.height        = height
	ap.windowCounter = 0
	// ap.CreateWindow("window")

	return ap
}

func (asp *aspirin)DebugPrint(msg string) {
	if asp.Debug {
		line := NewTextLine(msg)
		asp.debugWindow.ActivePane().AddLine(line, false)
	}
}


// print aspirin state for debugging
func Print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
       for _, c := range msg {
               termbox.SetCell(x, y, c, fg, bg)
               x++
       }
}

func Printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
       s := fmt.Sprintf(format, args...)
       Print_tb(x, y, fg, bg, s)
}

func Flush() {
       termbox.Flush()
}
