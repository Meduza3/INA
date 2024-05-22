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

//Każdy liść jest czarny. Jeśli węzeł jest czerwony, to obaj jego synowie są czarni
//Każda prosta ścieżka z ustalonego wezła do liścia ma tyle samo czarnych węzłów
