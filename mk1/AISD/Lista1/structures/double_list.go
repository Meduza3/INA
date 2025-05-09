package structures

type DoubleList struct {
	Length int
	First  *DoubleNode
	Last   *DoubleNode
}

type DoubleNode struct {
	Value    int
	Next     *DoubleNode
	Previous *DoubleNode
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		Length: 0,
		First:  nil,
		Last:   nil,
	}
}

func InsertDouble(l *DoubleList, i int) {
	newNode := &DoubleNode{Value: i, Next: nil}
	if l.First == nil {
		l.First = newNode
		l.Last = newNode
	} else {
		l.Last.Next = newNode
		newNode.Previous = l.Last
		l.Last = newNode
	}
	l.Length++
}

func MergeDouble(l1, l2 *DoubleList) {
	if l1.First == nil {
		*l1 = *l2
	} else if l2.First != nil {
		l1.Last.Next = l2.First
		l2.First.Previous = l1.Last
		l1.Last = l2.Last
		l1.Length += l2.Length
	}
}
