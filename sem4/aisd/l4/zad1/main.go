package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"trees/bst"
)

func main() {
	fmt.Println("hello!")
	tree := bst.NewEmpty()
	//Adding 50 random numbers and saving the array to s
	s := bst.AddNRandom(tree, 50, true)

	//Deleting 50 random numbers
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	for i := 0; i < 50; i++ {
		fmt.Printf("tree.Delete(%d)\n", s[i])
		tree.Delete(s[i])
		bst.PrintBST(tree.Root, tree.Height(), '-')
	}
	slices.Sort(s)
	//Adding 50 sorted numbers
	for i := 0; i < 50; i++ {
		fmt.Printf("tree.Insert(%d)\n", s[i])
		tree.Insert(s[i])
		bst.PrintBST(tree.Root, tree.Height(), '-')
	}

	//Removing 50 random numbers
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	for i := 0; i < 50; i++ {
		fmt.Printf("tree.Delete(%d)\n", s[i])
		tree.Delete(s[i])
		bst.PrintBST(tree.Root, tree.Height(), '-')
	}

}
