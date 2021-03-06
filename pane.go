package aspirin

import (
	"github.com/nsf/termbox-go"
	"github.com/gfx/go-visual_width"
	"strings"
)

type Pane interface {
	Init()

	setupEventLoop()
	ViewDidLoad()

	OnKey(func(ev Event))
	OnResize(func(ev Event))
	OnMouse(func(ev Event))
	OnError(func(ev Event))

	EventChannel() chan Event

	Update()

	setSize(int, int, int, int)
	Id() int
	setId(int) int
	Role() PaneRole
	setRole(PaneRole) PaneRole
	Left() Pane
	setLeft(Pane) Pane
	Right() Pane
	setRight(Pane) Pane
	Parent() Pane
	Size() *PaneSize
	setParent(Pane) Pane

	ContainWidth() int
	setContainWidth(int) int
	ContainHeight() int
	setContainHeight(int) int
	DivisionPoint() float32
	setDivisionPoint(float32) float32

	AddLine(l Line, setActive bool)
	ActiveLineIndex() int
	Lines() []Line
	setActiveLine(Line)
	MoveToLineIndex(lineIndex int)
	MoveNextLine()
	MovePrevLine()
	ScrollDown(size int)
	ScrollUp(size int)
}

// bbox: boundingboxとかにした方がよさそう，size，幅と高さだけっぽい
type PaneSize struct {
	x, y, width, height int
}

type PaneRole int
const (
	PRDisplay PaneRole = iota
	PRRoot
	PRVirticalSplit
	PRHorizontalSplit
)

type BasePane struct{
	id int
	size *PaneSize
	parent Pane
	left, right Pane
	role PaneRole

	activeLineIndex int
	topLineIndex int
	lines []Line

	onKey func(ev Event)
	onMouse func(ev Event)
	onResize func(ev Event)
	onError func(ev Event)
	eventChannel chan Event

	// for split pane
	containWidth  int
	containHeight int
	divisionPoint float32
}

func (bp *BasePane)Init() {
	// By default, no behavior are defined
	bp.onKey = (func(ev Event){})
	bp.onResize = (func(ev Event){})
	bp.onMouse = (func(ev Event){})
	bp.onError = (func(ev Event){})

	bp.eventChannel = make(chan Event)
	go bp.setupEventLoop()
}

func (bp *BasePane)Update() {
	paneSize := bp.size
 	var fgColor termbox.Attribute
 	var bgColor termbox.Attribute
	var line Line

	for index := 0; index < bp.size.height; index++ {
		if len(bp.lines) <= index {
			break
		}

		line = bp.lines[index + bp.topLineIndex]
		fgColor = line.Color().FgColor()
		bgColor = line.Color().BgColor()

		if bp.activeLineIndex == bp.findLine(line) {
			fgColor = line.ActiveColor().fgColor
			bgColor = line.ActiveColor().bgColor
		}

		text := visual_width.Truncate(line.Text(), true, paneSize.width, "...")
		// Clear brefor write
		Printf_tb(paneSize.x, paneSize.y + index, line.Color().FgColor(), line.Color().BgColor(), strings.Repeat(" ", paneSize.width))
		Printf_tb(paneSize.x, paneSize.y + index, fgColor, bgColor, text)
	}
	Flush()
}
func (bp *BasePane)ViewDidLoad() {
}
func (bp *BasePane)Id() int{
	return bp.id
}

func (bp *BasePane)Role() PaneRole{
	return bp.role
}
func (bp *BasePane)setRole(role PaneRole) PaneRole{
	bp.role = role
	return bp.role
}

func (bp *BasePane)setId(id int) int{
	bp.id = id
	return id
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
func (bp *BasePane)Size() *PaneSize{
	return bp.size
}
func (bp *BasePane)setSize(x, y, width, height int){
	bp.size = new(PaneSize)
	bp.size.x = x
	bp.size.y = y
	bp.size.width = width
	bp.size.height = height
}

// for split pane
func (bp *BasePane)ContainWidth() int{
	return bp.containWidth
}
func (bp *BasePane)setContainWidth(width int) int{
	bp.containWidth = width
	return bp.containWidth
}
func (bp *BasePane)ContainHeight() int{
	return bp.containHeight
}
func (bp *BasePane)setContainHeight(height int) int{
	bp.containHeight = height
	return bp.containHeight
}
func (bp *BasePane)DivisionPoint() float32{
	return bp.divisionPoint
}
func (bp *BasePane)setDivisionPoint(point float32) float32{
	bp.divisionPoint = point
	return bp.divisionPoint
}

func (bp *BasePane)AddLine(lo Line, setActive bool) {
	bp.lines = append(bp.lines, lo)

	if setActive {
		bp.setActiveLine(lo)
	}

	bp.Update()
}
func (bp *BasePane)Lines() []Line{
	return bp.lines
}
func (bp *BasePane)ActiveLineIndex() int{
	return bp.activeLineIndex
}
func (bp *BasePane)setActiveLine(lo Line){
	lineIndex := bp.findLine(lo)
	if lineIndex == -1 {
		// do nothing
		return
	}
	bp.MoveToLineIndex(lineIndex)
}
func (bp *BasePane)findLine(target Line) int{
	for index, line := range bp.lines {
		if line == target {
			return index
		}
	}
	return -1
}

func (bp *BasePane)MoveToLineIndex(lineIndex int){
	if lineIndex < bp.topLineIndex {
		diff := bp.topLineIndex - lineIndex
		bp.ScrollUp(diff)
		// heightからindexに直すので -1 してる
	} else if bp.topLineIndex + bp.size.height - 1 < lineIndex {
		diff := lineIndex - (bp.topLineIndex + bp.size.height - 1)
		bp.ScrollDown(diff)
	}

	bp.activeLineIndex = lineIndex
 	bp.Update()
}
func (bp *BasePane)MoveNextLine(){
	if bp.activeLineIndex == len(bp.lines) - 1 {
		return
	}

	bp.MoveToLineIndex(bp.activeLineIndex + 1)
}
func (bp *BasePane)MovePrevLine(){
	if bp.activeLineIndex == 0 {
		return
	}
	bp.MoveToLineIndex(bp.activeLineIndex - 1)
}
func (bp *BasePane)ScrollDown(size int){
	if (bp.topLineIndex + bp.size.height < len(bp.lines)) {
		bp.topLineIndex += 1
	}
}
func (bp *BasePane)ScrollUp(size int){
	if (bp.topLineIndex > 0) {
		bp.topLineIndex -= size
	} else {
		bp.topLineIndex = 0
	}
}
func (bp *BasePane)findLastLine(l Line) Line{
	// var lastLine Line
	// if l.Next() != nil {
	// 	lastLine = bp.findLastLine(l.Next())
	// } else {
	// 	lastLine = l
	// }
	return bp.lines[:1][0]
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
		case termbox.EventResize:
			go bp.onResize(ev)
		case termbox.EventMouse:
			go bp.onMouse(ev)
		case termbox.EventError:
			go bp.onError(ev)

		}
	}
}
