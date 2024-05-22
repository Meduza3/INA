package bst

type Tree struct {
	Root *Node
}

func New(root *Node) *Tree {
	tr := &Tree{}
	tr.Root = root
	return tr
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
}

func (t *Tree) insert(node *Node) {
	if t.Root == node {
		return
	}
	t.Root = t.Root.insert(node)
}

func (n *Node) insert(node *Node) *Node {
	if n == nil {
		return node
	}
	if n.Key < node.Key {
		n.Right = n.Right.insert(node)
	} else {
		n.Left = n.Left.insert(node)
	}
	return n
}

func main() {

}
