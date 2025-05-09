package rbt

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
