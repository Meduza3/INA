package rbt

import "fmt"

var leftTrace = make([]rune, 5000)
var rightTrace = make([]rune, 5000)

func PrintRBT(root *Node, depth int, prefix rune) {
	if root == nil {
		return
	}
	if root.Left != nil {
		PrintRBT(root.Left, depth+1, '/')
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
		PrintRBT(root.Right, depth+1, '\\')
	}
}
