package main

import (
	"fmt"

	"math/rand/v2"

	"./structures"
)

func main() {

	// Zadanie 1.
	stack := structures.Struct[int]{}
	queue := structures.Queue[int]{}

	for i := 0; i < 50; i++ {
		stack.Push(i)
		queue.Enqueue(i)
	}

	fmt.Println("Stack Pop:")
	for i := 0; i < 50; i++ {
		val, _ := stack.Pop()
		fmt.Print(val, " ")
	}
	fmt.Println()

	fmt.Println("Queue Dequeue:")
	for i := 0; i < 50; i++ {
		val, _ := queue.Dequeue()
		fmt.Print(val, " ")
	}
	fmt.Println()

	// Zadanie 2 i 3.
	const arraySize = 10000
	const searchCount = 1000
	const maxValue = 100000

	T := make([]int, arraySize)
	for i := range T {
		T[i] = rand.IntN(maxValue + 1)
	}

	list := structures.NewList()
	for _, v := range T {
		structures.Insert(list, v)
	}

	doubleList := structures.NewDoubleList()
	for _, v := range T {
		structures.InsertDouble(doubleList, v)
	}

	searchCostSingle := func(list *structures.List, target int) int {
		cost := 0
		current := list.First
		for current != nil {
			cost++
			if current.Value == target {
				break
			}
			current = current.Next
		}
		return cost
	}

	searchCostDouble := func(list *structures.DoubleList, target int) int {
		cost := 0
		current := list.First
		for current != nil {
			cost++
			if current.Value == target {
				break
			}
			current = current.Next
		}
		return cost
	}

	// Wyszukiwanie liczb z T (na pewno istniejących) w obu listach
	totalCostSingle1 := 0
	totalCostDouble1 := 0
	for i := 0; i < searchCount; i++ {
		target := T[rand.IntN(arraySize)]
		totalCostSingle1 += searchCostSingle(list, target)
		totalCostDouble1 += searchCostDouble(doubleList, target)
	}

	// Wyszukiwanie losowych liczb z przedziału I (mogą nie istnieć) w obu listach
	totalCostSingle2 := 0
	totalCostDouble2 := 0
	for i := 0; i < searchCount; i++ {
		target := rand.IntN(maxValue + 1)
		totalCostSingle2 += searchCostSingle(list, target)
		totalCostDouble2 += searchCostDouble(doubleList, target)
	}

	fmt.Printf("\nSingle List - Average search cost (existing numbers): %.2f comparisons\n", float64(totalCostSingle1)/float64(searchCount))
	fmt.Printf("Single List - Average search cost (random numbers): %.2f comparisons\n", float64(totalCostSingle2)/float64(searchCount))
	fmt.Printf("\nDouble List - Average search cost (existing numbers): %.2f comparisons\n", float64(totalCostDouble1)/float64(searchCount))
	fmt.Printf("Double List - Average search cost (random numbers): %.2f comparisons\n", float64(totalCostDouble2)/float64(searchCount))
}
