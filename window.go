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

	return w
}

func (win *window)GetRootPane() *pane{
	return win.rootPane
}

func (win *window)GetActivePane() *pane{
	return win.activePane
}

func (win *window)SetActivePane(id int) *pane{
	// concretePaneいがいがActivePaneになるの許可しない方がよさそう
	targetPane := findPaneById(win.rootPane, id)
	if(targetPane == nil){
		return nil
	}

	win.activePane = targetPane
	return win.activePane
}

func  (win *window)SetSize(width, height int) *window{
	// miniPane実装後はこうなるはず
	// win.width  = width - win.miniPane.height
	win.width  = width
	win.height = height
	return win
}

func (win *window)CloseyPane(paneId int){
}

func (win *window)MoveToNextPane(){
}

func (win *window)MoveToPrevPane(){
}

func (win *window)MoveToPane(targetId int) {
}

func (win *window)SplitPane(targetPane *pane, splitType SplitType) *pane{
	if (targetPane.paneType == RootPane) {
		// TODO: エラーどうしよ
		return nil
	}

	var vp *pane
	if (splitType == VirticalSplit) {
		vp = newPane(win.paneCounter, VirticalSplitPane)
	} else if (splitType == HorizontalSplit) {
		vp = newPane(win.paneCounter, HorizontalSplitPane)
	}
	win.activePane   = vp
	win.paneCounter += 1

	vp.parent         = targetPane.parent
	if (targetPane.parent.left.id == targetPane.id) {
		vp.parent.left  = vp
	} else {
		vp.parent.right = vp
	}
	targetPane.parent = vp
	vp.left           = targetPane
	vp.right          = newPane(win.paneCounter, ConcretePane)
	vp.right.parent   = vp

	win.activePane    = vp.right
	win.paneCounter += 1

	return vp.right
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
	if (targetPane.paneType == ConcretePane) {
		// TODO: エラーを返す???
		return nil
	}

	p :=  newPane(win.paneCounter, ConcretePane)
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

func findPaneById(targetPane *pane, id int) *pane {
	if (targetPane.id == id) {
		return targetPane
	}

	if (targetPane.left != nil){
		return findPaneById(targetPane.left, id)
	}
	if (targetPane.right != nil){
		return findPaneById(targetPane.right, id)
	}
	return nil
}
