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

func NewSplitPane(splitPaneId int,targetPane Pane, paneRole PaneRole, divisionPoint float32) (Pane, PaneSize, PaneSize){

	sp  := new(SplitPane)
	sp.containWidth  = targetPane.Size().width
	sp.containHeight =  targetPane.Size().height

	var leftPaneSize, rightPaneSize PaneSize

	sp.Init()
	sp.setId(splitPaneId)
	sp.setRole(paneRole)

	if (paneRole == PRVirticalSplit) {
		spX := targetPane.Size().x + int(float32(targetPane.Size().width) * divisionPoint)
		spY := targetPane.Size().y
		spWidth := 1
		spHeight := targetPane.Size().height

		sp.setSize(spX, spY, spWidth, spHeight)

		leftPaneSize.x       = targetPane.Size().x
		leftPaneSize.y       = targetPane.Size().y
		leftPaneSize.width   = int(float32(targetPane.Size().width) * divisionPoint)
		leftPaneSize.height  = targetPane.Size().height

		rightPaneSize.x      = spX + spWidth
		rightPaneSize.y      = spY
		rightPaneSize.width  = int(float32(targetPane.Size().width) * (1 - divisionPoint)) - spWidth
		rightPaneSize.height = targetPane.Size().height

		if (targetPane.Size().width % 2 == 1) {
			rightPaneSize.width += spWidth
		}
	} else if (paneRole == PRHorizontalSplit) {
		spX := targetPane.Size().x
		spY := targetPane.Size().y + int(float32(targetPane.Size().height) * divisionPoint)
		spWidth := targetPane.Size().width
		spHeight := 1

		sp.setSize(spX, spY, spWidth, spHeight)

		leftPaneSize.x       = targetPane.Size().x
		leftPaneSize.y       = targetPane.Size().y
		leftPaneSize.width   = targetPane.Size().width
		leftPaneSize.height  = int(float32(targetPane.Size().height) * divisionPoint)

		rightPaneSize.x      = spX
		rightPaneSize.y      = spY + spHeight
		rightPaneSize.width  = targetPane.Size().width
		rightPaneSize.height = int(float32(targetPane.Size().height) * (1 - divisionPoint)) - spHeight

		if (targetPane.Size().height % 2 == 1) {
			rightPaneSize.height += spHeight
		}
	}

	return sp, leftPaneSize, rightPaneSize
}

func CalcChildrenSize(target Pane, divisionPoint float32) (leftPaneSize, rightPaneSize *PaneSize){
	leftPaneSize  = new(PaneSize)
	rightPaneSize = new(PaneSize)

	return
}

func approximationDivisionPoint(lines int, divisionPoint float32) (approximatedDivisionPoint float32) {
	return
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
