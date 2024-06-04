package rbt

func (t *Tree) Successor(x *Node) *Node {
	if x == nil {
		return nil
	}

	if x.Right != nil {
		return t.Min()
	}

	y := x.Parent
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

func (tr Tree) Min() *Node {
	node := tr.Root
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}
