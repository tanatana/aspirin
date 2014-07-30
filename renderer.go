package aspirin

type Renderer interface {
	Rendering()
}

type defaultRenderer struct {
	targetPane *pane
}

func newDefaultRenderer(p *pane) *defaultRenderer{
	dr := new(defaultRenderer)
	dr.targetPane = p
	return dr
}

func (dr *defaultRenderer) Rendering(contents []string, line int) {

}
