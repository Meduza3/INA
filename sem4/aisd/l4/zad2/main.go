package main

import (
	"trees/bst"
)

func main() {

	stats := bst.ComplexityResults{0, 0, 0}

	DRZEWO_DUZEJ_CIPKI := bst.NewEmpty()
	for i := 0; i < 20; i++ {
		for n := 10_000; n < 100_000; n = n + 10_000 {
			bst.AddNRandom(DRZEWO_DUZEJ_CIPKI, n, false, &stats)
		}
	}
}
