package rbt

func (t *Tree) Delete(key int) *Node {
	z := t.Search(key)

	if z == nil {
		return nil
	}

	ret := &Node{nil, nil, nil, z.Key, z.Color}

	var y *Node
	var x *Node

	if z.Left == nil || z.Right == nil {
		y = z
	} else {
		y = t.Succesor(z)
	}

}
