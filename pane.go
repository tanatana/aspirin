package aspirin

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Pane interface {
	viewDidLoad()
	OnKey(func(ev Event))
	setupEventLoop()
	EventChannel() chan Event
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
	onKey func(ev Event)
	eventChannel chan Event
}

func (bp *BasePane)Init() {
	bp.onKey = (func(ev Event) {
		fmt.Printf("onKey@%s\n", "BasePane")
	})

	bp.eventChannel = make(chan Event)
	go bp.setupEventLoop()
}

func (bp *BasePane)SetSize(x, y, width, height int){
	bp.size.x = x
	bp.size.y = y
	bp.size.width = width
	bp.size.height = height
}


func (bp *BasePane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "BasePane")
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

func (bp *BasePane)OnKey(f func(ev Event)){
	bp.onKey = f
}

func (bp *BasePane)EventChannel() chan Event{
	return bp.eventChannel
}

func (bp *BasePane)setupEventLoop() {
	for {
		ev := <- bp.eventChannel
		fmt.Printf("%v\n", ev)
		switch ev.Type {
		case termbox.EventKey:
			go bp.onKey(ev)
		}
	}
}
