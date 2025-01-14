package main

import (
	"dijkstra/dimacs"
	"dijkstra/graphs"
	"encoding/csv"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	algos := []string{"dijkstra", "dial", "radix"}
	folderPath := flag.String("f", "", "Ścieżka do folderu z plikami .gr")
	outputDir := flag.String("o", "results", "Ścieżka do folderu na wyniki")
	flag.Parse()

	if *folderPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Create output directory if not exists
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		log.Fatalf("Nie można utworzyć folderu wynikowego: %v", err)
	}
	// Read immediate entries in *folderPath
	entries, err := os.ReadDir(*folderPath)
	if err != nil {
		log.Fatalf("Błąd podczas odczytu katalogu: %v", err)
	}

	// We'll use a WaitGroup to wait for all goroutines.
	var wg sync.WaitGroup

	// For each entry in *folderPath:
	for _, entry := range entries {
		if entry.IsDir() {
			// We found a subdirectory => process in separate goroutine.
			subdirPath := filepath.Join(*folderPath, entry.Name())

			wg.Add(1)
			go func(dirPath string) {
				defer wg.Done()
				if err := processDirectory(dirPath, *outputDir, algos); err != nil {
					log.Printf("Błąd podczas przetwarzania katalogu %s: %v\n", dirPath, err)
				}
			}(subdirPath)
		} else {
			// If the entry is a file (not a directory) and ends with .gr,
			// we could process it directly here if needed.
			// Or you can ignore top-level files, depending on your preference.
			if filepath.Ext(entry.Name()) == ".gr" {
				wg.Add(1)
				go func(filePath string) {
					defer wg.Done()
					runExperiment(filePath, *outputDir, algos)
				}(filepath.Join(*folderPath, entry.Name()))
			}
		}
	}

	// Wait for all goroutines to finish before exiting main.
	wg.Wait()
}

// processDirectory walks one directory and calls runExperiment on each .gr file.
func processDirectory(dirPath string, outputDir string, algos []string) error {
	return filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// If we can’t read the directory or something else, let’s log it.
			// Returning err will stop walking this directory entirely,
			// but you can also choose to return nil if you want to ignore the error.
			return err
		}

		// If it’s a directory, do nothing (just keep walking).
		if d.IsDir() {
			return nil
		}

		// If it’s a file and has .gr extension, run the experiment:
		if filepath.Ext(path) == ".gr" {
			runExperiment(path, outputDir, algos)
		}
		return nil
	})
}

func runExperiment(filePath, outputDir string, algos []string) {
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Could not stat file: %s, skipping.\n", filePath)
		saveSkippedOrFailureResult(filePath, outputDir, "ERROR", "Could not stat file")
		return
	}

	// Skip files larger than 80 MB
	const maxSizeBytes = 80 * 1024 * 1024
	if info.Size() > maxSizeBytes {
		fmt.Printf("Skipping %s because it exceeds %d bytes.\n", filePath, maxSizeBytes)
		saveSkippedOrFailureResult(filePath, outputDir, "SKIPPED", "File exceeds size limit")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Processing failed for file: %s due to: %v\n", filePath, r)
			saveSkippedOrFailureResult(filePath, outputDir, "ERROR", fmt.Sprintf("Panic: %v", r))
		}
	}()

	// Extract family name and prepare output file
	family := filepath.Base(filepath.Dir(filePath))
	fileName := filepath.Base(filePath)
	outputFilePath := filepath.Join(outputDir, fmt.Sprintf("%s_results.csv", family))

	fmt.Printf("Processing file: %s, Family: %s\n", filePath, family)

	dataFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Nie można otworzyć pliku z danymi: %v", err)
	}
	defer dataFile.Close()

	graph, err := dimacs.ParseFile(dataFile)
	if err != nil {
		log.Fatalf("Błąd parsowania pliku z danymi: %v", err)
	}

	numVertices := graph.Size

	// Prepare CSV writer
	var isNewFile bool
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		isNewFile = true
	}

	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Nie można utworzyć pliku CSV: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	if isNewFile {
		err = writer.Write([]string{"Algorithm", "Family", "File", "GraphSize", "SingleSourceTime", "AverageRandomSourceTime", "Status", "Message"})
		if err != nil {
			log.Fatalf("Błąd zapisu nagłówka CSV: %v", err)
		}
	}

	// Run algorithms and record results
	for _, algo := range algos {
		start := time.Now()
		_, err = graphs.RunSelectedAlgorithm(graph, 0, algo)
		if err != nil {
			log.Fatalf("Błąd działania algorytmu: %v", err)
		}
		duration := time.Since(start)

		rand.Seed(42)
		randomSources := rand.Perm(numVertices)[:5]
		var totalDuration time.Duration
		for _, source := range randomSources {
			start := time.Now()
			_, err := graphs.RunSelectedAlgorithm(graph, source, algo)
			if err != nil {
				log.Fatalf("Błąd działania algorytmu: %v", err)
			}
			totalDuration += time.Since(start)
		}
		averageDuration := totalDuration / 5

		err = writer.Write([]string{
			algo,
			family,
			fileName,
			fmt.Sprintf("%d", numVertices),
			fmt.Sprintf("%v", duration),
			fmt.Sprintf("%v", averageDuration),
			"SUCCESS",
			"",
		})
		if err != nil {
			log.Fatalf("Błąd zapisu wyników do CSV: %v", err)
		}
	}

	fmt.Printf("Results saved to: %s\n", outputFilePath)
}

func saveSkippedOrFailureResult(filePath, outputDir, status, message string) {
	family := filepath.Base(filepath.Dir(filePath))
	fileName := filepath.Base(filePath)
	outputFilePath := filepath.Join(outputDir, fmt.Sprintf("%s_results.csv", family))

	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Nie można utworzyć pliku CSV: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	err = writer.Write([]string{
		"",
		family,
		fileName,
		"0", // Graph size is unknown
		"0", // SingleSourceTime is zero
		"0", // AverageRandomSourceTime is zero
		status,
		message,
	})
	if err != nil {
		log.Fatalf("Błąd zapisu wyniku błędu do CSV: %v", err)
	}
	fmt.Printf("%s recorded for file: %s, reason: %s\n", status, filePath, message)
}
