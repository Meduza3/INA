package bst

func (tr *Tree) Delete(key int) {
	tr.Root.delete(key)
}

func (tr *Tree) DeleteStats(key int, stats *ComplexityResults) {
	tr.Root.deleteStats(key, stats)
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

func (n *Node) deleteStats(key int, stats *ComplexityResults) *Node {
	if n == nil {
		return nil
	}

	stats.Porownania++
	if key < n.Key {
		stats.OdczytyPodstawienia++
		n.Left = n.Left.delete(key)
	} else if key > n.Key {
		stats.Porownania++
		stats.OdczytyPodstawienia++
		n.Right = n.Right.delete(key)
	} else {
		stats.Porownania++
		if n.Left == nil && n.Right == nil {
			stats.OdczytyPodstawienia++
			n = nil
		} else if n.Left == nil {
			stats.Porownania++
			stats.OdczytyPodstawienia++
			n.Right.Parent = n.Parent
			stats.OdczytyPodstawienia++
			n = n.Right
		} else if n.Right == nil {
			stats.Porownania++
			stats.OdczytyPodstawienia++
			n.Left.Parent = n.Parent
			stats.OdczytyPodstawienia++
			n = n.Left
		} else {
			min := n.Right.minValueNode()
			stats.OdczytyPodstawienia++
			n.Key = min.Key
			stats.OdczytyPodstawienia++
			n.Right = n.Right.delete(min.Key)
			stats.Porownania++
			if n.Right != nil {
				stats.OdczytyPodstawienia++
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
