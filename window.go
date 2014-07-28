package aspirin

type window struct {
	id int
	activePane *pane
	title string
	width, height int
	rootPane *pane
	paneCounter int
}

func newWindow(id int, title string) *window{
	// TODO: get terminal's width & height
	// TODO: create mini-pane like emacs's mini-buffer
	w        := new(window)
	w.id      = id
	w.title   = title
	w.width   = 80
	w.height  = 24
	w.paneCounter = 0

	w.initializePaneTree()

	// w.rootPane = w.createRootPane()

	return w
}

func (win *window)GetRootPane() *pane{
	return win.rootPane
}

func (win *window)CloseyPane(paneId int){
}

func (win *window)MoveToNextPane(){
}

func (win *window)MoveToPrevPane(){
}

func (win *window)MoveToPane(targetId int) {
}

func (win *window)initializePaneTree() {
	rp := win.createRootPane()
	win.rootPane = rp
	win.createConcretePane(rp)
}

func (win *window)createRootPane() *pane{
	rp := newPane(win.paneCounter, RootPane)

	win.activePane   = rp
	win.paneCounter += 1
	return rp
}

func (win *window)createConcretePane(targetPane *pane) *pane{
	p := new(pane)
	p.parent = targetPane

	if (targetPane.left == nil || targetPane.paneType == RootPane) {
		targetPane.left = p
	} else {
		targetPane.right = p
	}

	win.activePane  = p
	win.paneCounter += 1
	return p
}

func (win *window)splitPane(targetPane *pane, splitType SplitType) *pane{
	return nil
}
