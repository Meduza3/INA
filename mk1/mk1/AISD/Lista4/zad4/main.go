package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"trees/rbt"
)

const (
	iterations int = 20
)

func main() {
	file, err := os.Create("results_random.csv")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"N", "Average Comparisons", "Average Substitutions", "Average Height", "Max Comparisons", "Max Substitutions", "Max Height"})

	for n := 10_000; n <= 100_000; n += 10_000 {
		avg_results := rbt.ComplexityResults{Porownania: 0, OdczytyPodstawienia: 0, Wysokosc: 0}
		max_results := rbt.ComplexityResults{Porownania: 0, OdczytyPodstawienia: 0, Wysokosc: 0}

		for i := 0; i < iterations; i++ {
			stats := rbt.ComplexityResults{Porownania: 0, OdczytyPodstawienia: 0, Wysokosc: 0}
			tree := rbt.NewEmpty()

			//elements := rbt.AddNSortedStats(tree, n, false, &stats)
			elements := rbt.AddNRandomStats(tree, n, false, &stats)

			rand.Shuffle(len(elements), func(i, j int) { elements[i], elements[j] = elements[j], elements[i] })

			for i := 0; i < len(elements); i++ {
				tree.DeleteStats(elements[i], &stats)
			}
			if stats.Porownania > max_results.Porownania {
				max_results.Porownania = stats.Porownania
			}
			if stats.OdczytyPodstawienia > max_results.OdczytyPodstawienia {
				max_results.OdczytyPodstawienia = stats.OdczytyPodstawienia
			}
			if stats.Wysokosc > max_results.Wysokosc {
				max_results.Wysokosc = stats.Wysokosc
			}
			avg_results.Add(stats)
			fmt.Printf("n: %d, porownania: %d, podstawienia: %d, wysokosc: %d\n", n, stats.Porownania, stats.OdczytyPodstawienia, stats.Wysokosc)
		}
		avg_results.Divide(iterations)
		fmt.Printf("ŚREDNIA DLA N: %d, porownania: %d, podstawienia: %d, wysokosc: %d\n", n, avg_results.Porownania, avg_results.OdczytyPodstawienia, avg_results.Wysokosc)
		fmt.Printf("MAX DLA N: %d, porownania: %d, podstawienia: %d, wysokosc: %d\n", n, max_results.Porownania, max_results.OdczytyPodstawienia, max_results.Wysokosc)
		writer.Write([]string{
			strconv.Itoa(n),
			strconv.Itoa(avg_results.Porownania),
			strconv.Itoa(avg_results.OdczytyPodstawienia),
			strconv.Itoa(avg_results.Wysokosc),
			strconv.Itoa(max_results.Porownania),
			strconv.Itoa(max_results.OdczytyPodstawienia),
			strconv.Itoa(max_results.Wysokosc),
		})
	}
}
