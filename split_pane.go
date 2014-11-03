package aspirin

import (
	"fmt"
)

type SplitPane struct {
	BasePane
	Type SplitType
	Text string
}

func NewSplitPane(splitPaneId int,targetPane Pane, splitType SplitType) (Pane, PaneSize, PaneSize){
	var sp Pane
	var leftPaneSize, rightPaneSize PaneSize

	if (splitType == SplitVirtical) {
		sp = new(SplitPane)
		// sp.setId(win.paneCounter)
		sp.setId(splitPaneId)
		sp.setSize(targetPane.Size().x + targetPane.Size().width/2,
			targetPane.Size().y,
			1,
			targetPane.Size().height)

		leftPaneSize.x       = targetPane.Size().x
		leftPaneSize.y       = targetPane.Size().y
		leftPaneSize.width   = targetPane.Size().width/2
		leftPaneSize.height  = targetPane.Size().height

		rightPaneSize.x      = targetPane.Size().x + targetPane.Size().width/2 + 1
		rightPaneSize.y      = targetPane.Size().y
		rightPaneSize.width  = targetPane.Size().width/2 - 1
		rightPaneSize.height = targetPane.Size().height

		if (targetPane.Size().width % 2 == 1) {
			rightPaneSize.width += 1
		}
	} else if (splitType == SplitHorizontal) {
		sp = new(SplitPane)
		sp.setId(splitPaneId)
		sp.setSize(targetPane.Size().x, targetPane.Size().y + targetPane.Size().height/2, targetPane.Size().width, 1)

		leftPaneSize.x       = targetPane.Size().x
		leftPaneSize.y       = targetPane.Size().y
		leftPaneSize.width   = targetPane.Size().width
		leftPaneSize.height  = targetPane.Size().height/2

		rightPaneSize.x      = targetPane.Size().x
		rightPaneSize.y      = targetPane.Size().y + targetPane.Size().height/2 + 1
		rightPaneSize.width  = targetPane.Size().width
		rightPaneSize.height = targetPane.Size().height/2 - 1

		if (targetPane.Size().height % 2 == 1) {
			rightPaneSize.height += 1
		}
	}
	return sp, leftPaneSize, rightPaneSize
}

func (rp *SplitPane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "SplitPane")
}
