package bst

func (tr *Tree) Delete(key int) {
	tr.Root.delete(key)
}

func (n *Node) delete(key int) *Node {
	if n == nil {
		return nil
	}

	if key < n.Key {
		n.Left = n.Left.delete(key)
	} else if key > n.Key {
		n.Right = n.Right.delete(key)
	} else {
		if n.Left == nil && n.Right == nil {
			n = nil
		} else if n.Left == nil {
			n.Right.Parent = n.Parent
			n = n.Right
		} else if n.Right == nil {
			n.Left.Parent = n.Parent
			n = n.Left
		} else {
			min := n.Right.minValueNode()
			n.Key = min.Key
			n.Right = n.Right.delete(min.Key)
			if n.Right != nil {
				n.Right.Parent = n
			}
		}
	}
	return n
}

func (n *Node) minValueNode() *Node {
	current := n

	for current.Left != nil {
		current = current.Left
	}
	return current
}
