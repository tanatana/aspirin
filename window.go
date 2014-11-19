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

func (w *window)MoveToNextPane() Pane{
	nextPane := findNextPane(w.activePane)

	if nextPane == nil {
		nextPane = findFirstPane(w.rootPane)
	}

	debugLine := NewTextLine(fmt.Sprintf("next pane id is %v", nextPane.Id()))
	w.activePane.AddLine(debugLine, false)

	w.activePane = nextPane
	return w.activePane
}
func (w *window)MoveToPrevPane() Pane{
	prevPane := findPrevPane(w.activePane)

	if prevPane == nil {
		prevPane = findLastPane(w.rootPane)
	}

	debugLine := NewTextLine(fmt.Sprintf("prev pane id is %v", prevPane.Id()))
	w.activePane.AddLine(debugLine, false)

	w.activePane = prevPane
	return w.activePane
}
func (w *window)MoveToFirstPane() Pane{
	firstPane := findFirstPane(w.rootPane)
	w.activePane = firstPane
	return w.activePane
}
func (w *window)MoveToLastPane() Pane{
	lastPane := findLastPane(w.rootPane)
	w.activePane = lastPane
	return w.activePane
}
func (w *window)MoveToPane() {}

func findNextPane(target Pane) Pane{
	parent := target.Parent()
	if parent.Parent() == nil {
		return nil
	}

	if parent.Left().Id() == target.Id() {
		// 自分が左にぶら下がってる状態で，兄弟の右が分割ペインなら
		// その右ペインを探しに行く
		// そうでなければ右を返す
		if parent.Right().Role() == PRVirticalSplit ||
			parent.Right().Role() == PRHorizontalSplit {
			return findNextPane(parent.Right())
		} else {
			return parent.Right()
		}
	}

	// 自分が右にぶら下がってる状況でのNextを探す場合は
	// いったん上に戻って右を探す
	if parent.Right().Id() == target.Id() {
		return findNextPane(parent)
	}

	panic("something wrong")
}
func findPrevPane(target Pane) Pane{
	parent := target.Parent()
	if parent.Role() == PRRoot {
		return nil
	}

	if parent.Right().Id() == target.Id() {
		if parent.Left().Role() == PRVirticalSplit ||
			parent.Left().Role() == PRHorizontalSplit {
			return findPrevPane(parent.Left())
		} else {
			return parent.Left()
		}
	}

	if parent.Left().Id() == target.Id() {
		return findPrevPane(parent)
	}
	panic("something wrong")
}

func findFirstPane(target Pane) Pane{
	if target.Left().Role() == PRVirticalSplit ||
		target.Left().Role() == PRHorizontalSplit {
		return findFirstPane(target.Left())
	}
	return target.Left()
}

func findLastPane(target Pane) Pane{
	if target.Right() != nil {
		if target.Right().Role() == PRVirticalSplit ||
			target.Right().Role() == PRHorizontalSplit {
			return findLastPane(target.Right())
		}
		return target.Right()
	} else if target.Left().Role() == PRVirticalSplit ||
		target.Left().Role() == PRHorizontalSplit {
		return findLastPane(target.Left())
	}
	return target.Left()

}

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

	newPane.viewDidLoad()

	return sp.Right()
}

func (w *window)SetInitialPane(child Pane, changeActivePane bool) {
	child.setId(w.paneCounter)
	child.setParent(w.rootPane)
	child.setSize(0, 0, w.Width(), w.Height())

	w.rootPane.setLeft(child)
	if (changeActivePane) {
		w.activePane = child
	}

	child.viewDidLoad()

	w.paneCounter += 1
}
