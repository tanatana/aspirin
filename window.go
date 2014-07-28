package aspirin

type window struct {
	id int
	activePane int
	title string
	width, height int
	rootPane *rootPane
	paneCounter int
}

func newWindow(id int, title string) *window{
	w        := new(window)
	w.id      = id
	w.activePane = -1
	w.title   = title
	w.width   = 80
	w.height  = 24
	w.paneCounter = 0
	w.rootPane = newRootPane()
	// w.rootPane = newRootPane()

	// cp := newConcretePane()
	// w.rootPane.SetFirstPane(cp)

	return w
}

func (win *window)GetRootPane() *rootPane{
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
