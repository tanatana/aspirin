package aspirin

import (
	"github.com/nsf/termbox-go"
)

type Pane interface {
	viewDidLoad()
	setupEventLoop()

	OnKey(func(ev Event))
	OnResize(func(ev Event))
	OnMouse(func(ev Event))
	OnError(func(ev Event))

	EventChannel() chan Event

	Update(x, y int, lo LineObject)

	SetSize(int, int, int, int)
	Id() int
	setId(int)
	Left() Pane
	setLeft(Pane) Pane
	Right() Pane
	setRight(Pane) Pane
	Parent() Pane
	Size() PaneSize
	setParent(Pane) Pane

	AddLineObject(LineObject)
}

// bbox: boundingboxとかにした方がよさそう，size，幅と高さだけっぽい
type PaneSize struct {
	x, y, width, height int
}

type SplitType int
const (
	VirticalSplit SplitType = iota
	HorizontalSplit
)


type BasePane struct{
	id int
	size PaneSize
	parent Pane
	left, right Pane

	rootLineObject LineObject
	activeLineObject LineObject

	onKey func(ev Event)
	onMouse func(ev Event)
	onResize func(ev Event)
	onError func(ev Event)
	eventChannel chan Event
}

func (bp *BasePane)Init() {
	// By default, no behavior are defined
	bp.onKey = (func(ev Event){})
	bp.onResize = (func(ev Event){})
	bp.onMouse = (func(ev Event){})
	bp.onError = (func(ev Event){})

	rootLineObj := newRootLineObject()
	bp.rootLineObject = rootLineObj
	bp.activeLineObject = rootLineObj

	bp.eventChannel = make(chan Event)
	go bp.setupEventLoop()
}

func (bp *BasePane)SetSize(x, y, width, height int){
	bp.size.x = x
	bp.size.y = y
	bp.size.width = width
	bp.size.height = height
}

func (bp *BasePane)Update(x, y int, lo LineObject) {
	fgColor := termbox.ColorDefault
	bgColor := termbox.ColorDefault

	if (lo == bp.activeLineObject) {
		fgColor = termbox.ColorWhite
		bgColor = termbox.ColorGreen
	}

	Printf_tb(x, y, fgColor, bgColor, lo.Text())

	if (lo.Next() != nil) {
		bp.Update(x, y + 1, lo.Next())
	} else {
		Printf_tb(70, y, fgColor, bgColor, "%d", x)
		Flush()
	}
}

func (bp *BasePane)viewDidLoad() {


}

func (bp *BasePane)Id() int{
	return bp.id
}
func (bp *BasePane)setId(id int){
	bp.id = id
}
func (bp *BasePane)Left() Pane{
	return bp.left
}
func (bp *BasePane)setLeft(p Pane) Pane{
	bp.left = p
	return bp.left
}
func (bp *BasePane)Right() Pane{
	return bp.right
}
func (bp *BasePane)setRight(p Pane) Pane{
	bp.right = p
	return bp.right
}
func (bp *BasePane)Parent() Pane{
	return bp.parent
}
func (bp *BasePane)setParent(p Pane) Pane{
	bp.parent = p
	return bp.parent
}
func (bp *BasePane)Size() PaneSize{
	return bp.size
}

func (bp *BasePane)AddLineObject(lo LineObject) {
	bp.activeLineObject.SetNext(lo)
	lo.SetPrev(bp.activeLineObject)
	bp.activeLineObject = lo

	bp.Update(0, 0, bp.rootLineObject)
}

func (bp *BasePane)OnKey(f func(ev Event)){
	bp.onKey = f
}
func (bp *BasePane)OnMouse(f func(ev Event)){
	bp.onMouse = f
}
func (bp *BasePane)OnResize(f func(ev Event)){
	bp.onResize = f
}
func (bp *BasePane)OnError(f func(ev Event)){
	bp.onError = f
}

func (bp *BasePane)EventChannel() chan Event{
	return bp.eventChannel
}

func (bp *BasePane)setupEventLoop() {
	for {
		ev := <- bp.eventChannel
		switch ev.Type {
		case termbox.EventKey:
			go bp.onKey(ev)
		}
	}
}
