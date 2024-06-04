package rbt

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

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
			PrintRBT(tree.Root, tree.Height(), '-')
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
			PrintRBT(tree.Root, tree.Height(), '-')
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
			PrintRBT(tree.Root, tree.Height(), '-')
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
