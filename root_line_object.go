package aspirin


type RootLineObject struct {
	LineObjectBase
}

func newRootLineObject() LineObject{
	lo := new(RootLineObject)

	return lo
}
