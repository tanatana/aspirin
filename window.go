package aspirin

type window struct {
	id int
	activePane int
	title string
	width, height int
	rootPane pane
}

func newWindow(id int) *window{
	w        := new(window)
	w.id      = id
	w.activePane = -1
	w.title   = "Aspirin"
	w.width   = 80
	w.height  = 24
	w.rootPane = newPane(3)

	return w
}

func (win *window)GetRootPane() pane{
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
