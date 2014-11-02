package aspirin

type LineObject interface {
	setupEventLoop()

	SetAction(func(ev Event))
	SetNext(LineObject)
	Next() LineObject
	SetPrev(LineObject)
	Prev() LineObject
	SetText(string)
	Text() string

	EventChannel() chan Event
}

type LineObjectBase struct {
	eventChannel chan Event
	action func(e Event)
	next LineObject
	prev LineObject
	text string
}

func (lob *LineObjectBase)SetAction(callback func(e Event)) {
	lob.action = callback
}


func (lob *LineObjectBase)SetNext(nextLineObj LineObject) {
	lob.next = nextLineObj
}
func (lob *LineObjectBase)Next() LineObject{
	return lob.next
}
func (lob *LineObjectBase)SetPrev(prevLineObj LineObject) {
	lob.prev = prevLineObj
}
func (lob *LineObjectBase)Prev() LineObject{
	return lob.prev
}
func (lob *LineObjectBase)SetText(newText string) {
	lob.text = newText
}
func (lob *LineObjectBase)Text() string{
	return lob.text
}


func (lob *LineObjectBase)setupEventLoop() {
	lob.eventChannel = make(chan Event)
}

func (lob *LineObjectBase)EventChannel() chan Event{
	return lob.eventChannel
}
