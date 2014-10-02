package aspirin

type window struct {
	id int
	activePane *pane
	title string
	width, height int
	rootPane *pane
	paneCounter int
}

func newWindow(id int, title string, width, height int) *window{
	// TODO: create mini-pane like emacs's mini-buffer
	w        := new(window)
	w.id      = id
	w.title   = title
	w.width   = width
	w.height  = height
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

func (win *window)ClosePane(target *pane) {
	parent := target.parent
	if (parent.paneType == RootPane) {
		// TODO: エラーどうしよ
		panic("can't split")
	}
	newParent := parent.parent
	keep := parent.left

	if keep.id == target.id {
		 keep = parent.right
	}

	// merge adjoining pane
	newPaneSize := calcMergedPaneSize(target, keep, parent)
	keep.setSize(newPaneSize.width, newPaneSize.height)
	keep.setPosition(newPaneSize.x, newPaneSize.y)
	if newParent.left.id == parent.id {
	    newParent.left = keep
	} else {
	    newParent.right = keep
	}

	keep.parent = newParent
	win.activePane = keep
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
		panic("can't split")
	}
	var sp *pane
	var leftPaneSize, rightPaneSize PaneSize

	if (splitType == VirticalSplit) {
		sp = newPane(win.paneCounter, VirticalSplitPane, targetPane.x + targetPane.width/2, targetPane.y, 1, targetPane.height)
		leftPaneSize.x      = targetPane.x
		leftPaneSize.y      = targetPane.y
		leftPaneSize.width  = targetPane.width/2 - 1
		leftPaneSize.height = targetPane.height

		rightPaneSize.x      = targetPane.x + targetPane.width/2 + 1
		rightPaneSize.y      = targetPane.y
		rightPaneSize.width  = targetPane.width/2 - 1
		rightPaneSize.height = targetPane.height

		if targetPane.width % 2 == 1 {
			rightPaneSize.width += 1
		}
	} else if (splitType == HorizontalSplit) {
		sp = newPane(win.paneCounter, HorizontalSplitPane, targetPane.x, targetPane.y + targetPane.height/2, targetPane.width, 1)
		leftPaneSize.x      = targetPane.x
		leftPaneSize.y      = targetPane.y
		leftPaneSize.width  = targetPane.width
		leftPaneSize.height = targetPane.height/2 - 1

		rightPaneSize.x      = targetPane.x
		rightPaneSize.y      = targetPane.y + targetPane.height/2 + 1
		rightPaneSize.width  = targetPane.width
		rightPaneSize.height = targetPane.height/2 - 1

		if targetPane.height % 2 == 1 {
			rightPaneSize.height += 1
		}

	}
	win.activePane   = sp
	win.paneCounter += 1

	sp.parent = targetPane.parent
	if (targetPane.parent.left.id == targetPane.id) {
		sp.parent.left  = sp
	} else {
		sp.parent.right = sp
	}
	targetPane.parent = sp
	sp.left           = targetPane
	sp.left.setSize(leftPaneSize.width, leftPaneSize.height);
	sp.left.setPosition(leftPaneSize.x, leftPaneSize.y);
	sp.right          = newPane(win.paneCounter, ConcretePane, 0, 0, win.width, win.height)
	sp.right.parent   = sp
	sp.right.setSize(rightPaneSize.width, rightPaneSize.height);
	sp.right.setPosition(rightPaneSize.x, rightPaneSize.y);
	win.activePane    = sp.right
	win.paneCounter += 1

	return sp.right
}


func (win *window)initializePaneTree() {
	rp := win.createRootPane()
	win.rootPane = rp
	win.createConcretePane(rp)
}

func (win *window)createRootPane() *pane{
	rp := newPane(win.paneCounter, RootPane, 0, 0, win.width, win.height)

	win.activePane   = rp
	win.paneCounter += 1
	return rp
}

func (win *window)createConcretePane(targetPane *pane) *pane{
	if (targetPane.paneType == ConcretePane) {
		// TODO: エラーを返す???
		return nil
	}

	p :=  newPane(win.paneCounter, ConcretePane, 0, 0, win.width, win.height)
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

func (win *window)refleshPaneTree(rootPane *pane) {
	rootPane.reflesh()
	if (rootPane.left != nil){
		win.refleshPaneTree(rootPane.left)
	}
	if (rootPane.right != nil){
		win.refleshPaneTree(rootPane.right)
	}
	return
}

func (win *window)refleshPane(targetPane *pane) {
	targetPane.reflesh()
}

func calcMergedPaneSize(target, bro, parent *pane) PaneSize{
	var size PaneSize
	if (target.x < bro.x) {
		size.x = target.x
	} else {
		size.x = bro.x
	}

	if (target.y < bro.y) {
		size.y = target.y
	} else {
		size.y = bro.y
	}


	if (target.x == bro.x) {
		size.width = target.width
		size.height = target.height + bro.height + parent.height
	} else {
		size.width = target.width + bro.width + parent.width
		size.height = target.height
	}

	return size
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
