package aspirin

import (
	"github.com/nsf/termbox-go"
)

type window struct {
	id int
	title string
	width, height int
	latestPaneId int // ペインを作成するたびに +1 される
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

	p := newRootPane(w.latestPaneId, width, height)
	w.rootPane = p
	w.activePane = p
	w.latestPaneId += 1

	go w.setupEventLoop()

	return w
}
func (w *window)Init() {
	go w.setupEventLoop()
}
func (w *window)OnKey(f func(ev Event)){
	w.onKey = f
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
	w.activePane = w.findNextPane()
	return w.activePane
}

func (w *window)findNextPane() Pane{
	// FIXME: pending
	pn := w.displayPaneList.Index(w.activePane)

	if pn.next == nil {
		return w.displayPaneList.FirstPane()
	}
	return pn.NextPane()
}

func (w *window)MoveToPrevPane() Pane{
	w.activePane = w.findPrevPane()
	return w.activePane
}
func (w *window)findPrevPane() Pane{
	// FIXME: pending
	pn := w.displayPaneList.Index(w.activePane)

	if pn.prev == nil {
		return w.displayPaneList.LastPane()
	}
	return pn.PrevPane()
}

func (w *window)MoveToFirstPane() Pane{
	w.activePane = w.displayPaneList.FirstPane()
	return w.activePane
}
func (w *window)MoveToLastPane()  Pane{
	w.activePane = w.displayPaneList.LastPane()
	return w.activePane
}

func (win *window)SplitPane(targetPane, newPane Pane, paneRole PaneRole) Pane{
	// if (targetPane.parent == nil) {
	// 	// TODO: エラーどうしよ
	// 	panic("can't split")
	// }

	sp, leftPaneSize, rightPaneSize := NewSplitPane(win.latestPaneId, targetPane, paneRole)
	win.latestPaneId += 1

	sp.setParent(targetPane.Parent())
	if (targetPane.Parent().Left().Id() == targetPane.Id()) {
		sp.Parent().setLeft(sp)
	} else {
		sp.Parent().setRight(sp)
	}

	sp.ViewDidLoad()

	targetPane.setParent(sp)

	sp.setLeft(targetPane)
	sp.Left().setSize(leftPaneSize.x, leftPaneSize.y, leftPaneSize.width, leftPaneSize.height);

	sp.setRight(newPane)
	sp.Right().setId(win.latestPaneId);
	sp.Right().setSize(0, 0, win.width, win.height)
	sp.Right().setParent(sp)
	sp.Right().setSize(rightPaneSize.x, rightPaneSize.y, rightPaneSize.width, rightPaneSize.height);

	win.activePane    = sp.Right()
	win.latestPaneId += 1

	win.refleshDisplayPaneList()
	newPane.ViewDidLoad()

	return sp.Right()
}

func (win *window)ClosePane(target Pane, movePrev bool) {
	// displayPaneListが1未満の場合ClosePane出来ない
	if win.displayPaneList.length <= 1 {
		return
	}

	var nextPane Pane
	if movePrev {
		nextPane = win.findPrevPane()
	} else {
		nextPane = win.findNextPane()
	}

	parent := target.Parent()
	var subTreeRoot Pane
	if parent.Left() == target {
		subTreeRoot = parent.Right()
	} else if parent.Right() == target {
		subTreeRoot = parent.Left()
	} else {
		panic("wtf!!!")
	}

	superParent := parent.Parent()
	if superParent.Right() == parent {
		superParent.setRight(subTreeRoot)
		subTreeRoot.setParent(superParent)
	} else if superParent.Left() == parent {
		superParent.setLeft(subTreeRoot)
		subTreeRoot.setParent(superParent)
	}

	win.refleshDisplayPaneList()
	win.activePane = nextPane

	// サイズの再設定
	// win.Resize()

	// アクティブペインを移す
	// activePane := newActivePane

	// 新しいサイズでの描画
	// win.Update()
}

func (w *window)SetInitialPane(child Pane) {
	child.setId(w.latestPaneId)
	child.setParent(w.rootPane)
	child.setSize(0, 0, w.Width(), w.Height())

	w.rootPane.setLeft(child)
	w.activePane = child
	child.ViewDidLoad()

	w.latestPaneId += 1
	w.refleshDisplayPaneList()
}
