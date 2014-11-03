package aspirin


type TextLineObject struct {
	LineObjectBase
}

func NewTextLineObject(text string) LineObject{
	tlo := new(TextLineObject)
	tlo.SetText(text)
	tlo.action = func(e Event){}
	return tlo
}

func (tlo *TextLineObject)SetAction(callback func(e Event)) {
	panic("I'm TextLineObject")
}
