package aspirin

import (
	"github.com/nsf/termbox-go"
	"fmt"
)

type window struct {
	id int
	title string
	width, height int
	paneCounter int
	activePane Pane
	rootPane Pane
	onKey func(ev Event)
	eventChannel chan Event
	displayPaneList *paneList
}

func (w *window)setupEventLoop() {
	for {
		ev := <- w.eventChannel
		switch ev.Type {
		case termbox.EventKey:
			go w.onKey(ev)
		}
		w.activePane.EventChannel() <- ev
	}
}

func NewWindow(title string, width, height int) *window{
	w             := new(window)
	w.title        = title
	w.width        = width
	w.height       = height

	w.eventChannel = make(chan Event)
	w.onKey = (func(e Event){})

	p := newRootPane(w.paneCounter, width, height)
	w.rootPane = p
	w.activePane = p
	w.paneCounter += 1

	go w.setupEventLoop()

	return w
}

func (w *window)RootPane() Pane{
	return w.rootPane
}

func (w *window)ActivePane() Pane{
	return w.activePane
}

func (w *window)Width() int{
	return w.width
}

func (w *window)Height() int{
	return w.height
}
func (w *window)refleshDisplayPaneList() {
	w.displayPaneList = new(paneList)
	w.updateDisplayPaneList(w.rootPane.Left())
}
func (w *window)updateDisplayPaneList(rootPane Pane) {
	if (rootPane.Left() != nil){
		w.updateDisplayPaneList(rootPane.Left())
	}
	if (rootPane.Right() != nil){
		w.updateDisplayPaneList(rootPane.Right())
	}

	if rootPane.Role() == PRDisplay {
		w.displayPaneList.Push(rootPane)
	}
	return
}
func (w *window)MoveToNextPane() Pane{
	// FIXME: pending
    debugLine := NewTextLine(fmt.Sprintf("current display length: %v", w.displayPaneList.length))
    w.activePane.AddLine(debugLine, false)

	return w.activePane
}
func (w *window)MoveToPrevPane() Pane{
	// FIXME: pending
    debugLine := NewTextLine(fmt.Sprintf("current display length: %v", w.displayPaneList.length))
    w.activePane.AddLine(debugLine, false)

	return w.activePane
}
// func (w *window)MoveToFirstPane() Pane{}
// func (w *window)MoveToLastPane()  Pane{}
// func (w *window)MoveToPane() {}

// func findNextPane (target Pane) Pane{}
// func findPrevPane (target Pane) Pane{}
// func findFirstPane(target Pane) Pane{}
// func findLastPane (target Pane) Pane{}

func (win *window)SplitPane(targetPane, newPane Pane, paneRole PaneRole) Pane{
	// if (targetPane.parent == nil) {
	// 	// TODO: エラーどうしよ
	// 	panic("can't split")
	// }

	sp, leftPaneSize, rightPaneSize := NewSplitPane(win.paneCounter, targetPane, paneRole)
	win.paneCounter += 1

	sp.setParent(targetPane.Parent())
	if (targetPane.Parent().Left().Id() == targetPane.Id()) {
		sp.Parent().setLeft(sp)
	} else {
		sp.Parent().setRight(sp)
	}

	sp.viewDidLoad()

	targetPane.setParent(sp)

	sp.setLeft(targetPane)
	sp.Left().setSize(leftPaneSize.x, leftPaneSize.y, leftPaneSize.width, leftPaneSize.height);

	sp.setRight(newPane)
	sp.Right().setId(win.paneCounter);
	sp.Right().setSize(0, 0, win.width, win.height)
	sp.Right().setParent(sp)
	sp.Right().setSize(rightPaneSize.x, rightPaneSize.y, rightPaneSize.width, rightPaneSize.height);

	win.activePane    = sp.Right()
	win.paneCounter += 1

	win.refleshDisplayPaneList()
	newPane.viewDidLoad()

	return sp.Right()
}

func (w *window)SetInitialPane(child Pane) {
	child.setId(w.paneCounter)
	child.setParent(w.rootPane)
	child.setSize(0, 0, w.Width(), w.Height())

	w.rootPane.setLeft(child)
	w.activePane = child
	child.viewDidLoad()

	w.paneCounter += 1
	w.refleshDisplayPaneList()
}
