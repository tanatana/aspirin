package aspirin

import (
//	"fmt"
)

type SplitPane struct {
	BasePane
}

func NewSplitPane(splitPaneId int,targetPane Pane, paneRole PaneRole) (Pane, PaneSize, PaneSize){
	// var sp Pane
	// sp = new(SplitPane)

	sp  := new(SplitPane)
	var leftPaneSize, rightPaneSize PaneSize

	sp.Init()
	sp.setId(splitPaneId)
	sp.setRole(paneRole)

	if (paneRole == PRVirticalSplit) {
		spX := targetPane.Size().x + targetPane.Size().width/2
		spY := targetPane.Size().y
		spWidth := 1
		spHeight := targetPane.Size().height

		sp.setSize(spX, spY, spWidth, spHeight)

		leftPaneSize.x       = targetPane.Size().x
		leftPaneSize.y       = targetPane.Size().y
		leftPaneSize.width   = targetPane.Size().width/2
		leftPaneSize.height  = targetPane.Size().height

		rightPaneSize.x      = spX + spWidth
		rightPaneSize.y      = spY
		rightPaneSize.width  = targetPane.Size().width/2 - spWidth
		rightPaneSize.height = targetPane.Size().height

		if (targetPane.Size().width % 2 == 1) {
			rightPaneSize.width += spWidth
		}
	} else if (paneRole == PRHorizontalSplit) {
		spX := targetPane.Size().x
		spY := targetPane.Size().y + targetPane.Size().height/2
		spWidth := targetPane.Size().width
		spHeight := 1

		sp.setSize(spX, spY, spWidth, spHeight)

		leftPaneSize.x       = targetPane.Size().x
		leftPaneSize.y       = targetPane.Size().y
		leftPaneSize.width   = targetPane.Size().width
		leftPaneSize.height  = targetPane.Size().height/2

		rightPaneSize.x      = spX
		rightPaneSize.y      = spY + spHeight
		rightPaneSize.width  = targetPane.Size().width
		rightPaneSize.height = targetPane.Size().height/2 - spHeight

		if (targetPane.Size().height % 2 == 1) {
			rightPaneSize.height += spHeight
		}
	}

	return sp, leftPaneSize, rightPaneSize
}

func (sp *SplitPane)viewDidLoad() {
	var splitLine Line
	if sp.role == PRHorizontalSplit{
		splitLine = NewTextLine("-")
		sp.AddLine(splitLine, false)
	} else if sp.role == PRVirticalSplit {
		splitLine = NewTextLine("|")
		sp.AddLine(splitLine, false)
	}
}
