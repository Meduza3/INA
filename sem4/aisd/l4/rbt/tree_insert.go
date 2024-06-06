package rbt

type Tree struct {
	Root *Node
}

type Color int

const (
	RED Color = iota
	BLACK
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
	Color  Color
}

func New(key int) *Tree {
	tr := &Tree{&Node{Key: key}}
	return tr
}

func NewEmpty() *Tree {
	tr := &Tree{}
	return tr
}

func (t *Tree) Height() int {
	return t.Root.height()
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}
	leftHeight := n.Left.height()
	rightHeight := n.Right.height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (t *Tree) Insert(key int) *Node {
	z := &Node{nil, nil, nil, key, BLACK}
	x := t.Root
	var y *Node = nil

	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else if x.Key < z.Key {
			x = x.Right
		} else {
			return x
		}
	}

	z.Parent = y
	if y == nil {
		t.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}

	t.insertFixup(z)
	return z
}

func (t *Tree) InsertStats(key int, stats *ComplexityResults) *Node {
	z := &Node{nil, nil, nil, key, BLACK}
	stats.OdczytyPodstawienia++
	x := t.Root
	var y *Node = nil

	for x != nil {
		y = x
		stats.Porownania++
		if z.Key < x.Key {
			stats.Porownania++
			stats.OdczytyPodstawienia++
			x = x.Left
		} else if x.Key < z.Key {
			stats.Porownania++
			stats.Porownania++
			stats.OdczytyPodstawienia++
			x = x.Right
		} else {
			stats.Porownania++
			stats.Porownania++
			return x
		}
	}

	stats.OdczytyPodstawienia++
	z.Parent = y
	if y == nil {
		stats.Porownania++
		stats.OdczytyPodstawienia++
		t.Root = z
	} else if z.Key < y.Key {
		stats.Porownania++
		stats.Porownania++
		stats.OdczytyPodstawienia++
		y.Left = z
	} else {
		stats.Porownania++
		stats.Porownania++
		stats.OdczytyPodstawienia++
		y.Right = z
	}

	t.insertFixup(z)
	return z
}

func (t *Tree) insertFixup(z *Node) {
	if z == nil {
		return
	}
	for z != t.Root && z.Parent.Color == RED {
		var G *Node // grandparent
		if z.Parent.Parent != nil {
			G = z.Parent.Parent
		} else {
			break // If there is no grandparent, break out of the loop
		}

		if z.Parent == G.Left {
			U := G.Right // uncle
			if U != nil && U.Color == RED {
				z.Parent.Color = BLACK
				U.Color = BLACK
				G.Color = RED
				z = G
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.leftRotate(z)
				}
				z.Parent.Color = BLACK
				G.Color = RED
				t.rightRotate(G)
			}
		} else {
			U := G.Left // uncle
			if U != nil && U.Color == RED {
				z.Parent.Color = BLACK
				U.Color = BLACK
				G.Color = RED
				z = G
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.rightRotate(z)
				}
				z.Parent.Color = BLACK
				G.Color = RED
				t.leftRotate(G)
			}
		}
	}
	t.Root.Color = BLACK // Ensure the root is always black
}

func (t *Tree) insertFixupStats(z *Node, stats *ComplexityResults) {

	if z.Parent == nil || z.Parent.Color == BLACK {
		z.Color = RED
		return
	}

	for z.Parent != nil && z.Parent.Color == RED {
		stats.Porownania++
		var G *Node // grandparent
		stats.Porownania++
		if z.Parent.Parent == nil {
			G = nil
		} else {
			stats.OdczytyPodstawienia++
			G = z.Parent.Parent
		}
		stats.Porownania++
		if z.Parent == G.Left {
			stats.OdczytyPodstawienia++
			U := G.Right // uncle
			stats.Porownania++
			if U != nil && U.Color == RED { // Case I2
				z.Parent.Color = BLACK
				U.Color = BLACK
				G.Color = RED
				stats.OdczytyPodstawienia++
				z = G
			} else {
				stats.Porownania++
				if z == z.Parent.Right { // Case I5
					stats.OdczytyPodstawienia++
					z = z.Parent
					t.leftRotateStats(z, stats)
				}
				// Case I6
				z.Parent.Color = BLACK
				G.Color = RED
				t.rightRotateStats(G, stats)
			}
		} else {
			stats.OdczytyPodstawienia++
			U := G.Left // uncle
			stats.Porownania++
			if U != nil && U.Color == RED { // Case I2
				z.Parent.Color = BLACK
				U.Color = BLACK
				G.Color = RED
				stats.OdczytyPodstawienia++
				z = G
			} else {
				stats.Porownania++
				if z == z.Parent.Left { // Case I5
					stats.OdczytyPodstawienia++
					z = z.Parent
					t.rightRotateStats(z, stats)
				}
				// Case I6
				z.Parent.Color = BLACK
				G.Color = RED
				t.leftRotateStats(G, stats)
			}
		}
	}
	t.Root.Color = BLACK // Case I4
}
