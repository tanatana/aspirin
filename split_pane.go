package aspirin

import (
	"fmt"
)

type SplitPane struct {
	BasePane
	Type SplitType
}

func (rp *SplitPane)viewDidLoad() {
	fmt.Printf("viewDidLoad@%s\n", "SplitPane")
}
