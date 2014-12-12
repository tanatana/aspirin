package aspirin

type Line interface {
	setupEventLoop()

	SetAction(func(ev Event))
	RunAction(Event)
	SetText(string)
	Text() string

	EventChannel() chan Event
}

type LineBase struct {
	eventChannel chan Event
	action func(e Event)
	text string
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


func (lob *LineBase)setupEventLoop() {
	lob.eventChannel = make(chan Event)
}

func (lob *LineBase)EventChannel() chan Event{
	return lob.eventChannel
}
