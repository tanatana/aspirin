package aspirin

import (
	"github.com/nsf/termbox-go"
)

type Color struct {
	fgColor termbox.Attribute
	bgColor termbox.Attribute
}

func NewColor(fg, bg termbox.Attribute) *Color{
	c := new(Color)
	c.fgColor = fg
	c.bgColor = bg
	return c
}

func (c *Color)FgColor() termbox.Attribute{
	return c.fgColor
}

func (c *Color)BgColor() termbox.Attribute{
	return c.bgColor
}
