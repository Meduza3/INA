package main

import (
	"fmt"
	"math/rand"
	"slices"
	"trees/rbt"
)

func main() {
	tree := rbt.NewEmpty()

	s := rbt.AddNRandom(tree, 50, true)
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	for i := 0; i < 50; i++ {
		fmt.Printf("tree.Delete(%d)\n", s[i])
		tree.Delete(s[i])
		rbt.PrintRBT(tree.Root, tree.Height(), '-')
	}
	slices.Sort(s)
	//Adding 50 sorted numbers
	for i := 0; i < 50; i++ {
		fmt.Printf("tree.Insert(%d)\n", s[i])
		tree.Insert(s[i])
		rbt.PrintRBT(tree.Root, tree.Height(), '-')
	}

	//Removing 50 random numbers
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	for i := 0; i < 50; i++ {
		fmt.Printf("tree.Delete(%d)\n", s[i])
		tree.Delete(s[i])
		rbt.PrintRBT(tree.Root, tree.Height(), '-')
	}
}
