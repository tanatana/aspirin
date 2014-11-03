package aspirin


type TextLine struct {
	LineBase
}

func NewTextLine(text string) Line{
	tlo := new(TextLine)
	tlo.SetText(text)
	tlo.action = func(e Event){}
	return tlo
}

func (tlo *TextLine)SetAction(callback func(e Event)) {
	panic("I'm TextLine")
}
