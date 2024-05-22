package rbt

func (t *Tree) leftRotate(x *Node) {
	//
	//          |                                  |
	//          X                                  Y
	//         / \         left rotate            / \
	//        α  Y       ------------->         X   γ
	//          / \                           / \
	//         β  γ                          α   β
	//

	if x.Right == nil {
		return
	}

	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent

	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y

}

func (t *Tree) leftRotateStats(x *Node, stats *ComplexityResults) {
	//
	//          |                                  |
	//          X                                  Y
	//         / \         left rotate            / \
	//        α  Y       ------------->         X   γ
	//          / \                           / \
	//         β  γ                          α   β
	//

	stats.Porownania++
	if x.Right == nil {
		return
	}

	y := x.Right
	stats.OdczytyPodstawienia++
	x.Right = y.Left
	stats.Porownania++
	if y.Left != nil {
		stats.OdczytyPodstawienia++
		y.Left.Parent = x
	}
	stats.OdczytyPodstawienia++
	y.Parent = x.Parent
	stats.Porownania++
	if x.Parent == nil {
		stats.OdczytyPodstawienia++
		t.Root = y
	} else if x == x.Parent.Left {
		stats.Porownania++
		stats.OdczytyPodstawienia++
		x.Parent.Left = y
	} else {
		stats.Porownania++
		stats.Porownania++
		stats.OdczytyPodstawienia++
		x.Parent.Right = y
	}
	stats.OdczytyPodstawienia++
	stats.OdczytyPodstawienia++
	y.Left = x
	x.Parent = y

}

func (t *Tree) rightRotate(x *Node) {
	if x.Left == nil {
		return
	}

	//
	//          |                                  |
	//          X                                  Y
	//         / \         right rotate           / \
	//        Y   γ      ------------->         α  X
	//       / \                                    / \
	//      α  β                                 β  γ

	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent

	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Right = x
	x.Parent = y
}

func (t *Tree) rightRotateStats(x *Node, stats *ComplexityResults) {

	//
	//          |                                  |
	//          X                                  Y
	//         / \         right rotate           / \
	//        Y   γ      ------------->         α  X
	//       / \                                    / \
	//      α  β                                 β  γ

	stats.Porownania++
	if x.Left == nil {
		return
	}
	y := x.Left
	stats.OdczytyPodstawienia++
	x.Left = y.Right
	stats.Porownania++
	if y.Right != nil {
		stats.OdczytyPodstawienia++
		y.Right.Parent = x
	}
	stats.OdczytyPodstawienia++
	y.Parent = x.Parent
	stats.Porownania++
	if x.Parent == nil {
		stats.OdczytyPodstawienia++
		t.Root = y
	} else if x == x.Parent.Left {
		stats.Porownania++
		stats.OdczytyPodstawienia++
		x.Parent.Left = y
	} else {
		stats.Porownania++
		stats.Porownania++
		stats.OdczytyPodstawienia++
		x.Parent.Right = y
	}
	stats.OdczytyPodstawienia++
	stats.OdczytyPodstawienia++
	y.Right = x
	x.Parent = y
}
