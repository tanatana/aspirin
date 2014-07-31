package aspirin

type Renderable interface {
	rendering()
}

type defaultRenderer struct {
	targetPane *pane
}

func newDefaultRenderer(p *pane) *defaultRenderer{
	dr := new(defaultRenderer)
	dr.targetPane = p
	return dr
}

func (dr *defaultRenderer) rendering(contents []string, line int) []string{
	var view []string
	for _, line := range contents {
		view = append(view, line)
	}

	return view
}

type HelloWorldRenderer struct {
	targetPane *pane
}

func newHelloWorldRenderer(p *pane) *HelloWorldRenderer{
	hr := new(HelloWorldRenderer)
	hr.targetPane = p
	return hr
}

func (hr *HelloWorldRenderer) rendering(contents []string, line int) []string{
	var view []string
	view = append (view, "Hello World!")

	return view
}
