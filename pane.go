package aspirin

import (
	"fmt"
)

type Pane interface {
	viewDidLoad()
}

type paneBase struct{
	pane bool
	Id int
}

func (pb paneBase)viewDidLoad() {
	fmt.Printf("viewDidLoad from %s\n", "Pane::Base")
}
