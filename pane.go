package aspirin

type PaneType int

const (
	RootPane PaneType = iota
	VirticalSplitPane
	HorizontalSplitPane
	Pane
)

func (i PaneType) String() string {
    switch i {
    case RootPane:
        return "RootPane"
    case VirticalSplitPane:
        return "VirticalSplitPAne"
    case HorizontalSplitPane:
        return "HorizontalSplitPane"
    }
    return ""
}

type pane interface {
	GetPaneType() PaneType
}

func newPane(paneType PaneType) pane{
	return new(rootPane)
}

type rootPane struct {
	paneType PaneType

}

func (p *rootPane)GetPaneType() PaneType{
	return p.paneType
}

func newRootPane() pane{
	p := new(rootPane)
	p.paneType = RootPane
	return p
}
