package structures

type List struct {
	Length int
	First  *Node
}

type Node struct {
	Value int
	Next  *Node
}

func NewList() *List {
	return &List{
		Length: 0,
		First:  nil,
	}
}

func Insert(l *List, i int) {
	newNode := &Node{Value: i, Next: nil}
	if l.First == nil {
		l.First = newNode
	} else {
		next := l.First
		for next.Next != nil {
			next = next.Next
		}
		next.Next = newNode
	}
	l.Length++
}

func Merge(l1, l2 *List) {
	if l1.First == nil {
		l1.First = l2.First
	} else {
		next := l1.First
		for next.Next != nil {
			next = next.Next
		}
		next.Next = l2.First
	}
	l1.Length += l2.Length
}
