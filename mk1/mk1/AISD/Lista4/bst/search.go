package bst

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

func (tr Tree) Max() *Node {
	node := tr.Root
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tr Tree) Search(key int) *Node {
	node := tr.Root
	for node != nil {
		switch {
		case node.Key < key:
			node = node.Right
		case node.Key > key:
			node = node.Left
		default:
			return node
		}
	}
	return nil
}
