package rbt

func (t *Tree) Delete(key int) {
	var child *Node
	node := t.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if nodeColor(node) == BLACK {
			node.Color = nodeColor(child)
			t.deleteCase1(node)
		}
		t.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.Color = BLACK
		}
	}
}

func (t *Tree) DeleteStats(key int, stats *ComplexityResults) {
	var child *Node
	node := t.lookupStats(key, stats)
	stats.Porownania++
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNodeStats(stats)
		node.Key = pred.Key
		node = pred
		stats.OdczytyPodstawienia++
	}
	stats.Porownania++
	stats.Porownania++
	if node.Left == nil || node.Right == nil {
		stats.Porownania++
		if node.Right == nil {
			child = node.Left
			stats.OdczytyPodstawienia++
		} else {
			child = node.Right
			stats.OdczytyPodstawienia++
		}
		stats.Porownania++
		if nodeColor(node) == BLACK {
			node.Color = nodeColor(child)
			t.deleteCase1Stats(node, stats)
		}
		t.replaceNodeStats(node, child, stats)
		stats.Porownania++
		stats.Porownania++
		if node.Parent == nil && child != nil {
			child.Color = BLACK
		}
	}
}

func (tree *Tree) lookup(key int) *Node {
	node := tree.Root
	for node != nil {
		if key == node.Key {
			return node
		} else if key < node.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}

func (tree *Tree) lookupStats(key int, stats *ComplexityResults) *Node {
	node := tree.Root
	for node != nil {
		if key == node.Key {
			stats.Porownania++
			return node
		} else if key < node.Key {
			stats.Porownania++
			stats.Porownania++
			node = node.Left
			stats.OdczytyPodstawienia++
		} else {
			stats.Porownania++
			stats.Porownania++
			node = node.Right
			stats.OdczytyPodstawienia++
		}
	}
	return nil
}

func (tree *Tree) deleteCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *Tree) deleteCase1Stats(node *Node, stats *ComplexityResults) {
	stats.Porownania++
	if node.Parent == nil {
		return
	}
	tree.deleteCase2Stats(node, stats)
}

func (tree *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == RED {
		node.Parent.Color = RED
		sibling.Color = BLACK
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree) deleteCase2Stats(node *Node, stats *ComplexityResults) {
	sibling := node.sibling()
	if nodeColor(sibling) == RED {
		node.Parent.Color = RED
		sibling.Color = BLACK
		stats.Porownania++
		if node == node.Parent.Left {
			tree.rotateLeftStats(node.Parent, stats)
		} else {
			tree.rotateRightStats(node.Parent, stats)
		}
	}
	tree.deleteCase3Stats(node, stats)
}

func (tree *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == BLACK &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree) deleteCase3Stats(node *Node, stats *ComplexityResults) {
	sibling := node.sibling()
	if sibling != nil && nodeColor(node.Parent) == BLACK &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		tree.deleteCase1Stats(node.Parent, stats)
	} else {
		tree.deleteCase4Stats(node, stats)
	}
}

func (tree *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if sibling != nil && nodeColor(node.Parent) == RED &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		node.Parent.Color = BLACK
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree) deleteCase4Stats(node *Node, stats *ComplexityResults) {
	sibling := node.sibling()
	if sibling != nil && nodeColor(node.Parent) == RED &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == BLACK &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		node.Parent.Color = BLACK
	} else {
		tree.deleteCase5Stats(node, stats)
	}
}

func (tree *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if sibling != nil && node == node.Parent.Left &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == RED &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		sibling.Left.Color = BLACK
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == BLACK &&
		nodeColor(sibling.Left) == RED &&
		nodeColor(sibling.Right) == BLACK {
		sibling.Color = RED
		sibling.Right.Color = BLACK
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree) deleteCase5Stats(node *Node, stats *ComplexityResults) {
	sibling := node.sibling()
	// Ensure sibling and its children are not nil before accessing their properties
	if sibling != nil && sibling.Left != nil && sibling.Right != nil {
		if node == node.Parent.Left &&
			nodeColor(sibling) == BLACK &&
			nodeColor(sibling.Left) == RED &&
			nodeColor(sibling.Right) == BLACK {
			sibling.Color = RED
			sibling.Left.Color = BLACK
			tree.rotateRightStats(sibling, stats)
		} else if node == node.Parent.Right &&
			nodeColor(sibling) == BLACK &&
			nodeColor(sibling.Left) == RED &&
			nodeColor(sibling.Right) == BLACK {
			sibling.Color = RED
			sibling.Right.Color = BLACK
			tree.rotateLeftStats(sibling, stats)
		}
	}
	tree.deleteCase6Stats(node, stats)
}

func (tree *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	if sibling != nil {
		sibling.Color = node.Parent.Color
	}
	node.Parent.Color = BLACK
	if node == node.Parent.Left && nodeColor(sibling.Right) == RED {
		sibling.Right.Color = BLACK
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == RED {
		sibling.Left.Color = BLACK
		tree.rotateRight(node.Parent)
	}
}

func (tree *Tree) deleteCase6Stats(node *Node, stats *ComplexityResults) {
	sibling := node.sibling()
	if sibling != nil { // Ensure sibling is not nil before accessing its properties
		sibling.Color = node.Parent.Color
		node.Parent.Color = BLACK
		stats.Porownania++
		stats.Porownania++
		if node == node.Parent.Left && nodeColor(sibling.Right) == RED {
			if sibling.Right != nil { // Check if sibling.Right is not nil before accessing its Color
				sibling.Right.Color = BLACK
				tree.rotateLeftStats(node.Parent, stats)
			}
		} else if sibling.Left != nil && nodeColor(sibling.Left) == RED { // Check if sibling.Left is not nil
			sibling.Left.Color = BLACK
			tree.rotateRightStats(node.Parent, stats)
		}
	}
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (tree *Tree) rotateLeft(node *Node) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *Tree) rotateRight(node *Node) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *Tree) rotateLeftStats(node *Node, stats *ComplexityResults) {
	right := node.Right
	tree.replaceNodeStats(node, right, stats)
	node.Right = right.Left
	stats.OdczytyPodstawienia++
	stats.Porownania++
	if right.Left != nil {
		stats.OdczytyPodstawienia++
		right.Left.Parent = node
	}
	stats.OdczytyPodstawienia++
	stats.OdczytyPodstawienia++
	right.Left = node
	node.Parent = right
}

func (tree *Tree) rotateRightStats(node *Node, stats *ComplexityResults) {
	left := node.Left
	tree.replaceNodeStats(node, left, stats)
	node.Left = left.Right
	stats.OdczytyPodstawienia++
	stats.Porownania++
	if left.Right != nil {
		stats.OdczytyPodstawienia++
		left.Right.Parent = node
	}
	stats.OdczytyPodstawienia++
	stats.OdczytyPodstawienia++
	left.Right = node
	node.Parent = left
}

func (tree *Tree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *Tree) replaceNodeStats(old *Node, new *Node, stats *ComplexityResults) {
	stats.Porownania++
	if old.Parent == nil {
		tree.Root = new
	} else {
		stats.Porownania++
		stats.OdczytyPodstawienia++
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	stats.Porownania++
	if new != nil {
		stats.OdczytyPodstawienia++
		new.Parent = old.Parent
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (node *Node) maximumNodeStats(stats *ComplexityResults) *Node {
	stats.Porownania++
	if node == nil {
		return nil
	}
	stats.Porownania++
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func nodeColor(node *Node) Color {
	if node == nil {
		return BLACK
	}
	return node.Color
}
