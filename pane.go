package aspirin

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Pane interface {
	viewDidLoad()
	onKey(ev Event)
	setupEventLoop()
	EventChannel() chan Event
	SetSize(int, int, int, int)
	Left() Pane
	setLeft(Pane) Pane
	Right() Pane
	setRight(Pane) Pane
	Parent() Pane
	setParent(Pane) Pane
	Split(SplitType)
	Close()
}

type BasePane struct{
	id int
	x, y int
	width, height int
	parent Pane
	left, right Pane
	eventChannel chan Event
}

type SplitType int

const (
	VirticalSplit SplitType = iota
	HorizontalSplit
)

func (bp *BasePane)Init() {
	bp.eventChannel = make(chan Event)
	go bp.setupEventLoop()
}

func (bp *BasePane)SetSize(x, y, width, height int){
	bp.x = x
	bp.y = y
	bp.width = width
	bp.height = height
}


func (bp *BasePane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "BasePane")
}

func (bp *BasePane)onKey(ev Event) {
	fmt.Printf("onKey@%s\n", "BasePane")
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

func (bp *BasePane)Split(t SplitType) {}

func (bp *BasePane)Close() {}

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
