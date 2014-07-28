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

type SplitType int

const (
	VirticalSplit SplitType = iota
	HorizontalSplit
)

func (i SplitType) String() string {
    switch i {
    case VirticalSplit:
        return "VirticalSplit"
    case HorizontalSplit:
        return "HorizontalSplit"
    }
    return ""
}


type pane struct{
	id int
	paneType PaneType
	parent *pane
	left, right *pane
	x, y int
	width, height int
}

func newPane(id int, paneType PaneType) *pane{
	return new(pane)
}
