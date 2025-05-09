package main

import (
	"bufio"
	"fmt"
	"os"
)

func binarySearchRec(arr []int, low, high, target int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	switch {
	case arr[mid] == target:
		return true
	case arr[mid] > target:
		return binarySearchRec(arr, low, mid-1, target)
	default:
		return binarySearchRec(arr, mid+1, high, target)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, v int

	// --- If your input is "n v" on the first line, uncomment this and comment out the next Scan:
	// if _, err := fmt.Fscan(reader, &n, &v); err != nil {
	//     fmt.Fprintln(os.Stderr, "read n and v:", err)
	//     os.Exit(1)
	// }

	// Read n:
	if _, err := fmt.Fscan(reader, &n); err != nil {
		fmt.Fprintln(os.Stderr, "read n:", err)
		os.Exit(1)
	}

	// Read the n array elements:
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			fmt.Fprintln(os.Stderr, "read arr[", i, "]:", err)
			os.Exit(1)
		}
	}

	// Read v (the target):
	if _, err := fmt.Fscan(reader, &v); err != nil {
		fmt.Fprintln(os.Stderr, "read v:", err)
		os.Exit(1)
	}

	// Perform the search and print 1 if found, 0 otherwise:
	if binarySearchRec(arr, 0, n-1, v) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
