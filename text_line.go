package aspirin

import (
	"github.com/nsf/termbox-go"
)

type TextLine struct {
	LineBase
}

func NewTextLine(text string) Line{
	tlo := new(TextLine)
	tlo.SetText(text)
	tlo.action = func(e Event){}

	tlo.color = NewColor(termbox.ColorDefault, termbox.ColorDefault)
	// activeColor変わるの，actionがあるlineだけにしたいんだけど，
	// いまの実装でそうするとスクロール時に違和感が出るので
	// とりあえず色付けてる
	tlo.activeColor = NewColor(termbox.ColorDefault, termbox.ColorDefault)
	tlo.activeColor = NewColor(termbox.ColorWhite, termbox.ColorGreenq)

	return tlo
}

func (tlo *TextLine)SetAction(callback func(e Event)) {
	panic("I'm TextLine")
}
