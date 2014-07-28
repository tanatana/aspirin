package aspirin

type PaneType int

const (
	RootPane PaneType = iota
	VirticalSplitPane
	HorizontalSplitPane
	ConcretePane
)

func (i PaneType) String() string {
    switch i {
    case RootPane:
        return "RootPane"
    case VirticalSplitPane:
        return "VirticalSplitPane"
    case HorizontalSplitPane:
        return "HorizontalSplitPane"
	case ConcretePane:
		return "ConcretePane"
    }
    return ""
}

type rootPane struct {
	id int
	paneType PaneType
	firstPane *pane
}

func newRootPane() *rootPane{
	rp := new(rootPane)
	rp.id = 0
	rp.paneType = RootPane
	rp.firstPane = newPane()
	return rp
}

type pane struct{
	id int
	paneType PaneType
	parent *pane
	left, right *pane
	x, y int
	width, height int
}

func newPane() *pane{
	return new(pane)
}


// type pane interface {
// 	GetPaneType() PaneType
// }

// type rootPane struct {
// 	paneType PaneType
// 	fisrt pane
// }

// func (p *rootPane)GetPaneType() PaneType{
// 	return p.paneType
// }

// func (rp *rootPane)SetFirstPane(p pane) {
// 	rp.firstPane = p
// }

// func newRootPane() pane{
// 	p := new(rootPane)
// 	p.paneType = RootPane
// 	p.fisrt = newConcretePane()
// 	return p
// }

// type virticalSplitPane struct {
// 	paneType PaneType
// 	parent pane
// 	fisrt pane
// 	second pane
// }

// func (vp *virticalSplitPane)GetPaneType() PaneType{
// 	return vp.paneType
// }

// func newVirticalSplitPane(p pane) pane{
// 	vp := new(virticalSplitPane)
// 	vp.paneType = VirticalSplitPane
// 	vp.fisrt = p
// 	vp.second = newConcretePane()
// 	return vp
// }

// type horizontalSplitPane struct {
// 	paneType PaneType
// 	parent pane
// 	fisrt pane
// 	second pane
// }

// func (hp *horizontalSplitPane)GetPaneType() PaneType{
// 	return hp.paneType
// }

// func newHorizontalSplitPane(p pane) pane{
// 	hp := new(horizontalSplitPane)
// 	hp.paneType = RootPane
// 	hp.fisrt = p
// 	hp.second = newConcretePane()
// 	return hp
// }

// type concretePane struct {
// 	paneType PaneType
// 	parent pane
// }

// func (cp *concretePane)GetPaneType() PaneType{
// 	return cp.paneType
// }

// func newConcretePane() pane{
// 	hp := new(horizontalSplitPane)
// 	hp.paneType = RootPane
// 	return hp
// }
