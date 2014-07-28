package aspirin

type Aspirin struct {
	activeWindow int
	windows []*window
	windowCounter int
}

func NewAspirin() *Aspirin {
	newAspirin := new(Aspirin)
	newAspirin.windowCounter = 0
	newAspirin.CreateWindow("window")
	return newAspirin
}

func (ap *Aspirin)CreateWindow(title string) {
	w := newWindow(ap.windowCounter, title)
	ap.windows = append(ap.windows, w)
	ap.activeWindow = w.id
	ap.windowCounter += 1

}

func (ap *Aspirin)GetWindows() []*window{
	return ap.windows
}

func (ap *Aspirin)ActiveWindow() *window{
	return ap.windows[ap.activeWindow]
}
