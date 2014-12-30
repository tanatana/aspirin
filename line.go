package aspirin

type Line interface {
	setupEventLoop()

	SetAction(func(ev Event))
	RunAction(Event)
	SetText(string)
	Text() string
	SetColor(*Color)
	Color() *Color
	SetActiveColor(*Color)
	ActiveColor() *Color

	EventChannel() chan Event
}

type LineBase struct {
	eventChannel chan Event
	action func(e Event)
	text string
	color *Color
	activeColor *Color
}

func (lob *LineBase)SetAction(callback func(e Event)) {
	lob.action = callback
}

func (lob *LineBase)RunAction(e Event)  {
	lob.action(e)
}

func (lob *LineBase)SetText(newText string) {
	lob.text = newText
}
func (lob *LineBase)Text() string{
	return lob.text
}
func (lob *LineBase)SetColor(c *Color){
	lob.color = c
}
func (lob *LineBase)Color() *Color{
	return lob.color
}
func (lob *LineBase)SetActiveColor(c *Color){
	lob.activeColor = c
}
func (lob *LineBase)ActiveColor() *Color{
	return lob.activeColor
}


func (lob *LineBase)setupEventLoop() {
	lob.eventChannel = make(chan Event)
}

func (lob *LineBase)EventChannel() chan Event{
	return lob.eventChannel
}
