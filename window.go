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
	p.setSize(0, 0, width, height)
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
	nextPane := w.findNextPane(w.activePane)
	if nextPane == nil {
		nextPane = w.findFirstPane(w.rootPane)
	}
	w.activePane = nextPane

	return w.activePane
}
func (w *window)MoveToPrevPane() {}
func (w *window)MoveToPane() {}

func (w *window)findNextPane(target Pane) Pane{
	parent := target.Parent()
	if parent.Parent() == nil {
		return nil
	}

	if parent.Left().Id() == target.Id() {
		// 昔はPaneにpaneTypeってのあって比較できたけどいまは
		// rowPainを匿名フィールドに持つPainインターフェースを
		// 実装したオブジェクトってだけなので見極めるすべがない
		// 枝の状態から推定は出来るけど

		// 自分が左にぶら下がってる状態で，兄弟の右が分割ペインなら
		// その右ペインを探しに行く
		// そうでなければ右を返す
		// if parent.Right().paneType == VirticalSplitPane ||
		// 	parent.Right().paneType == HorizontalSplitPane {
		// 	return findPrevPane(parent.Right())
		// } else {
		// 	return parent.Right()
		// }
	}

	// 自分が右にぶら下がってる状況でのNextを探す場合は
	// いったん上に戻って右を探す
	if parent.Right().Id() == target.Id() {
		return w.findNextPane(parent)
	}

	panic("something wrong")
}

func (w *window)findFirstPane(target Pane) Pane{
	// findNextPaneと同様の問題がある
	// if target.Left().paneType == VirticalSplitPane ||
	// 	target.Left().paneType == HorizontalSplitPane {
	// 	return findFirstPane(target.Left())
	// }
	return target.Left()
}
func (w *window)findLastPane(p Pane) {}

func (win *window)SplitPane(targetPane, newPane Pane, splitType SplitType) Pane{
	// if (targetPane.parent == nil) {
	// 	// TODO: エラーどうしよ
	// 	panic("can't split")
	// }
	sp, leftPaneSize, rightPaneSize := NewSplitPane(win.paneCounter, targetPane, splitType)
	// win.activePane   = sp

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
