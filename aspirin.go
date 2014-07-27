package aspirin

type Aspirin struct {
	activeWindow int
	windows []*window
}

func NewAspirin() *Aspirin {
	newAspirin := new(Aspirin)
	newAspirin.CreateWindow()
	return newAspirin
}

func (ap *Aspirin)CreateWindow() {
	// lenではなく最後のwindowsのidを使うように修正する
	w := newWindow(len(ap.windows))
	ap.windows = append(ap.windows, w)
	ap.activeWindow = w.id
}

func (ap *Aspirin)GetWindows() []*window{
	return ap.windows
}

func (ap *Aspirin)ActiveWindow() *window{
	return ap.windows[ap.activeWindow]
}
