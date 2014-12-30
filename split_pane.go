package aspirin

import (
	"strings"
)

type SplitPane struct {
	BasePane
}

func NewSplitPane(splitPaneId int,targetPane Pane, splitPaneRole PaneRole, divisionPoint float32) (Pane, *PaneSize, *PaneSize){
	sp := new(SplitPane)
	targetSize := targetPane.Size()

	sp.Init()
	sp.setId(splitPaneId)
	sp.setRole(splitPaneRole)

	sp.setContainWidth(targetSize.width)
	sp.setContainHeight(targetSize.height)
	sp.setDivisionPoint(divisionPoint)

	spSize, lpSize, rpSize := calcSplitSize(targetSize, splitPaneRole, divisionPoint)
	sp.setSize(spSize.x, spSize.y, spSize.width, spSize.height)

	return sp, lpSize, rpSize
}

func calcSplitSize(targetSize *PaneSize,  splitPaneRole PaneRole, divisionPoint float32) (spSize, leftSize, rightSize *PaneSize){

	spSize    = new(PaneSize)
	leftSize  = new(PaneSize)
	rightSize = new(PaneSize)

	if (splitPaneRole == PRVirticalSplit) {
		spSize.x      = targetSize.x + int(float32(targetSize.width) * divisionPoint)
		spSize.y      = targetSize.y
		spSize.width  = 1
		spSize.height = targetSize.height

		leftSize.x       = targetSize.x
		leftSize.y       = targetSize.y
		leftSize.width   = int(float32(targetSize.width) * divisionPoint)
		leftSize.height  = targetSize.height

		rightSize.x      = spSize.x + spSize.width
		rightSize.y      = spSize.y
		rightSize.width  = int(float32(targetSize.width) * (1 - divisionPoint)) - spSize.width
		rightSize.height = targetSize.height

		if (targetSize.width % 2 == 1) {
			rightSize.width += spSize.width
		}
	} else if (splitPaneRole == PRHorizontalSplit) {
		spSize.x      = targetSize.x
		spSize.y      = targetSize.y + int(float32(targetSize.height) * divisionPoint)
		spSize.width  = targetSize.width
		spSize.height = 1

		leftSize.x       = targetSize.x
		leftSize.y       = targetSize.y
		leftSize.width   = targetSize.width
		leftSize.height  = int(float32(targetSize.height) * divisionPoint)

		rightSize.x      = spSize.x
		rightSize.y      = spSize.y + spSize.height
		rightSize.width  = targetSize.width
		rightSize.height = int(float32(targetSize.height) * (1 - divisionPoint)) - spSize.height

		if (targetSize.height % 2 == 1) {
			rightSize.height += spSize.height
		}
	}

	return
}

func (sp *SplitPane)ViewDidLoad() {
	var splitLine Line
	if sp.role == PRHorizontalSplit{
		splitLine = NewTextLine(strings.Repeat("-", sp.Size().width))
		// splitLine = NewTextLine(fmt.Sprintf("%v", sp))
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
