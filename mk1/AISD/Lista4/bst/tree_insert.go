package bst

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

type Tree struct {
	Root *Node
}

func New(key int) *Tree {
	tr := &Tree{&Node{Key: key}}
	return tr
}

func NewEmpty() *Tree {
	tr := &Tree{}
	return tr
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
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

func (t *Tree) Insert(key int) {
	node := &Node{Key: key}
	if t.Root == node {
		return
	}
	t.Root = t.Root.insert(node)
}

func (t *Tree) InsertStats(key int, stats *ComplexityResults) {
	node := &Node{Key: key}
	if t.Root == node {
		stats.Porownania++
		return
	}
	stats.OdczytyPodstawienia++
	t.Root = t.Root.insertStats(node, stats)
}

func (n *Node) insert(node *Node) *Node {
	if n == nil {
		return node
	}
	if n.Key < node.Key {
		node.Parent = n
		n.Right = n.Right.insert(node)
	} else {
		node.Parent = n
		n.Left = n.Left.insert(node)
	}
	return n
}

func (n *Node) insertStats(node *Node, stats *ComplexityResults) *Node {
	if n == nil {
		return node
	}
	stats.Porownania++
	if n.Key < node.Key {
		stats.OdczytyPodstawienia++
		node.Parent = n
		stats.OdczytyPodstawienia++
		n.Right = n.Right.insert(node)
	} else {
		stats.OdczytyPodstawienia++
		node.Parent = n
		stats.OdczytyPodstawienia++
		n.Left = n.Left.insert(node)
	}
	return n
}

func AddNRandom(tree *Tree, n int, print bool) []int {
	s := make([]int, n)
	for i := 0; i < 50; i++ {
		random := rand.IntN(100)
		if !contains(s, random) {
			s[i] = random
		} else {
			i--
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("tree.Insert(%d)\n", s[i])
		tree.Insert(s[i])
		if print {
			PrintBST(tree.Root, tree.Height(), '-')
		}
	}
	return s
}

func AddNRandomStats(tree *Tree, n int, print bool, stats *ComplexityResults) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		random := rand.IntN(2 * n)
		if !contains(s, random) {
			s[i] = random
		} else {
			i--
		}
	}

	for i := 0; i < n; i++ {
		tree.InsertStats(s[i], stats)
		if print {
			fmt.Printf("tree.Insert(%d)\n", s[i])
			PrintBST(tree.Root, tree.Height(), '-')
		}
	}
	stats.Wysokosc = tree.Height()
	return s
}

func AddNSortedStats(tree *Tree, n int, print bool, stats *ComplexityResults) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		random := rand.IntN(2 * n)
		if !contains(s, random) {
			s[i] = random
		} else {
			i--
		}
	}

	slices.Sort(s)

	for i := 0; i < n; i++ {
		tree.InsertStats(s[i], stats)
		if print {
			fmt.Printf("tree.Insert(%d)\n", s[i])
			PrintBST(tree.Root, tree.Height(), '-')
		}
	}
	stats.Wysokosc = tree.Height()
	return s
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

var leftTrace = make([]rune, 5000)
var rightTrace = make([]rune, 5000)

func PrintBST(root *Node, depth int, prefix rune) {
	if root == nil {
		return
	}
	if root.Left != nil {
		PrintBST(root.Left, depth+1, '/')
	}
	if prefix == '/' {
		leftTrace[depth-1] = '|'
	}
	if prefix == '\\' {
		rightTrace[depth-1] = ' '
	}
	if depth == 0 {
		fmt.Print("-")
	}
	if depth > 0 {
		fmt.Print(" ")
	}
	for i := 0; i < depth-1; i++ {
		if leftTrace[i] == '|' || rightTrace[i] == '|' {
			fmt.Print("| ")
		} else {
			fmt.Print("  ")
		}
	}
	if depth > 0 {
		fmt.Printf("%c-", prefix)
	}
	fmt.Printf("[%d]\n", root.Key)
	leftTrace[depth] = ' '
	if root.Right != nil {
		rightTrace[depth] = '|'
		PrintBST(root.Right, depth+1, '\\')
	}
}
