package aspirin


type RootLine struct {
	LineBase
}

func newRootLine() Line{
	lo := new(RootLine)

	return lo
}
