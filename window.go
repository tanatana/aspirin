package aspirin

import (
	"github.com/nsf/termbox-go"
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

	p := new(RootPane)
	p.Init()
	p.setId(w.paneCounter)
	p.SetSize(0, 0, width, height)
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

func (w *window)MoveToNextPane() {}
func (w *window)MoveToPrevPane() {}
func (w *window)MoveToPane() {}

func (win *window)SplitPane(targetPane, newPane Pane, splitType SplitType) Pane{
	// if (targetPane.parent == nil) {
	// 	// TODO: エラーどうしよ
	// 	panic("can't split")
	// }
	var sp Pane
	var leftPaneSize, rightPaneSize PaneSize

	// if (splitType == VirticalSplit) {
	// 	sp = newPane(win.paneCounter, VirticalSplitPane, targetPane.x + targetPane.width/2, targetPane.y, 1, targetPane.height)

	// 	leftPaneSize.x       = targetPane.x
	// 	leftPaneSize.y       = targetPane.y
	// 	leftPaneSize.width   = targetPane.width/2
	// 	leftPaneSize.height  = targetPane.height

	// 	rightPaneSize.x      = targetPane.x + targetPane.width/2 + 1
	// 	rightPaneSize.y      = targetPane.y
	// 	rightPaneSize.width  = targetPane.width/2 - 1
	// 	rightPaneSize.height = targetPane.height

	// 	if (targetPane.width % 2 == 1) {
	// 		rightPaneSize.width += 1
	// 	}


	// } else if (splitType == HorizontalSplit) {
	sp = new(SplitPane)
	sp.setId(win.paneCounter)
	sp.SetSize(targetPane.Size().x, targetPane.Size().y + targetPane.Size().height/2, targetPane.Size().width, 1)

	leftPaneSize.x       = targetPane.Size().x
	leftPaneSize.y       = targetPane.Size().y
	leftPaneSize.width   = targetPane.Size().width
	leftPaneSize.height  = targetPane.Size().height/2

	rightPaneSize.x      = targetPane.Size().x
	rightPaneSize.y      = targetPane.Size().y + targetPane.Size().height/2 + 1
	rightPaneSize.width  = targetPane.Size().width
	rightPaneSize.height = targetPane.Size().height/2 - 1

	if (targetPane.Size().height % 2 == 1) {
		rightPaneSize.height += 1
	}
	// }

	win.activePane   = sp

	sp.setParent(targetPane.Parent())
	if (targetPane.Parent().Left().Id() == targetPane.Id()) {
		sp.Parent().setLeft(sp)
	} else {
		sp.Parent().setRight(sp)
	}

	targetPane.setParent(sp)
	sp.setLeft(targetPane)
	sp.Left().SetSize(leftPaneSize.x, leftPaneSize.y, leftPaneSize.width, leftPaneSize.height);
	//	sp.setRight(newPane(win.paneCounter, ConcretePane, 0, 0, win.width, win.height))
	sp.setRight(newPane)
	sp.Right().setId(win.paneCounter);
	sp.Right().SetSize(0, 0, win.width, win.height)
	sp.Right().setParent(sp)
	sp.Right().SetSize(rightPaneSize.x, rightPaneSize.y, rightPaneSize.width, rightPaneSize.height);
	win.activePane    = sp.Right()
	win.paneCounter += 1

	return sp.Right()
}

func (w *window)SetInitialPane(child Pane, changeActivePane bool) {
	child.setId(w.paneCounter)
	child.setParent(w.rootPane)
	w.rootPane.setLeft(child)
	if (changeActivePane) {
		w.activePane = child
	}
	w.paneCounter += 1
}
