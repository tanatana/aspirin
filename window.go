package aspirin

type Window struct {
	id int
	activePane Pane
	title string
	width, height int
	rootPane Pane
}

func newWindow(id int, title string, width, height int) *Window{
	// TODO: create mini-pane like emacs's mini-buffer
	w        := new(Window)
	w.id      = id
	w.title   = title
	w.width   = width
	w.height  = height
	w.rootPane  = new(paneBase)
	// w.initializePaneTree()

	return w
}
