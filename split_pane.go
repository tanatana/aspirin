package aspirin

import (
//	"fmt"
	"strings"
)

type SplitPane struct {
	BasePane
	containWidth int
	containHeight int
	position float32
}

func NewSplitPane(splitPaneId int,targetPane Pane, paneRole PaneRole) (Pane, PaneSize, PaneSize){

	sp  := new(SplitPane)
	sp.containWidth  = targetPane.Size().width
	sp.containHeight =  targetPane.Size().height

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

func (sp *SplitPane)ViewDidLoad() {
	var splitLine Line
	if sp.role == PRHorizontalSplit{
		splitLine = NewTextLine(strings.Repeat("-", sp.Size().width))
		splitLine.SetActiveColor(splitLine.Color())
		sp.AddLine(splitLine, true)
	} else if sp.role == PRVirticalSplit {
		for i := 0; i < sp.Size().height; i++ {
			splitLine = NewTextLine("|")
			splitLine.SetActiveColor(splitLine.Color())
			sp.AddLine(splitLine, true)
		}
	}
}
