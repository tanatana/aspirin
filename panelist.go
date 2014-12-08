package aspirin

type paneList struct{
	first *paneNode
	last *paneNode
	length int
}

type paneNode struct{
	pl *paneList
	self Pane
	next *paneNode
	prev *paneNode
}

func (pl *paneList)FirstPane() Pane{
	return pl.first.self
}
func (pl *paneList)LastPane() Pane{
	return pl.last.self
}
func (pl *paneList)Index(targetPane Pane) *paneNode{
	return pl._Index(pl.first, targetPane)
}

func (pl *paneList)_Index(targetNode *paneNode, targetPane Pane) *paneNode{
	if targetNode.self == targetPane {
		return targetNode
	}

	if targetNode.next != nil {
		return pl._Index(targetNode.next, targetPane)
	}

	return nil
}

func (pl *paneList)Push(p Pane) *paneList{
	// ary = [1,2,3]
	// ary.push(4) #=> [1,2,3,4]
	pn := newPaneNode(p, pl)
	if pl.first == nil && pl.last == nil {
		pl.first = pn
		pl.last  = pn
	} else {
		pl.last.next = pn
		pn.prev = pl.last
		pl.last = pn
	}

	pl.length += 1
	return pl
}

func (pl *paneList)Unshift(p Pane) *paneList{
	// ary = [1,2,3]
	// ary.unshift(0) #=> [0,1,2,3]
	pn := newPaneNode(p, pl)
	pl.first.prev = pn
	pn.next = pl.first
	pl.first = pn

	pl.length += 1
	return pl
}
func (pl *paneList)Pop(p Pane) Pane{
	// ary = [1,2,3,4]
	// ary.pop() #=> 4
	// ary       #=> [1,2,3]
	var oldLast *paneNode
	oldLast = pl.last
	pl.last = pl.last.prev

	pl.length -= 1
	return oldLast.self
}
func (pl *paneList)Shift(Pane) Pane{
	// ary = [1,2,3,4]
	// ary.shift() #=> 1
	// ary         #=> [2,3,4]
	var oldFirst *paneNode
	if pl.first.next != nil {
		oldFirst = pl.first
		pl.first = pl.first.next
	}

	pl.length -= 1
	return oldFirst.self
}


func newPaneNode(p Pane, pl *paneList) *paneNode{
	pn := new(paneNode)
	pn.self = p
	pn.pl   = pl
	return pn
}
func (pn *paneNode)Next() *paneNode{
	if pn.next != nil {
		return pn.next
	}

	return nil
}
func (pn *paneNode)Prev() *paneNode{
	if pn.prev != nil {
		return pn.prev
	}

	return nil
}
func (pn *paneNode)NextPane() Pane{
	if pn.next != nil {
		return pn.next.self
	}

	return nil
}
func (pn *paneNode)PrevPane() Pane{
	if pn.prev != nil {
		return pn.prev.self
	}

	return nil
}
