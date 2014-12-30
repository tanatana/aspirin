package aspirin

type ActionLine struct {
	LineBase
}

func NewActionLine(text string, color, actionColor *Color, action func(e Event)) Line{
	alo := new(ActionLine)
	alo.SetText(text)
	alo.action = func(e Event){}

	alo.color = color
	alo.activeColor = actionColor

	alo.action = action

	return alo
}

func (alo *ActionLine)SetAction(action func(e Event)) {
	alo.action = action
}
