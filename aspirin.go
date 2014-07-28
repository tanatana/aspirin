package aspirin

import("fmt")

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

func (ap *Aspirin)CreateWindow(title string) *window{
	w := newWindow(ap.windowCounter, title)
	ap.windows = append(ap.windows, w)
	ap.activeWindow = w.id
	ap.windowCounter += 1
	return w
}

func (ap *Aspirin)GetWindows() []*window{
	return ap.windows
}

func (ap *Aspirin)GetActiveWindow() *window{
	return ap.windows[ap.activeWindow]
}

func DrawPaneTree(targetPane *pane, floor int) {
	fmt.Print("\t\t")
	for i := 0; i < floor; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("%v\n", *targetPane)
	if (targetPane.left != nil){
		DrawPaneTree(targetPane.left, floor + 1)
	}

	if (targetPane.right != nil){
		DrawPaneTree(targetPane.right, floor + 1)
	}

}
