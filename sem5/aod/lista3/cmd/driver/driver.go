package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func main() {
	folderPath := "../ch9-1.1/inputs" // Replace with your folder path
	outputDir := "results"            // Replace with your output directory

	// Read immediate entries in folderPath
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var crashedFamilies []string

	// For each subdirectory or file, spawn a separate process
	for _, entry := range entries {
		entryPath := filepath.Join(folderPath, entry.Name())
		if entry.IsDir() || filepath.Ext(entry.Name()) == ".gr" {
			wg.Add(1)
			go func(path, family string) {
				defer wg.Done()
				if err := runExperimentProcess(path, outputDir); err != nil {
					mu.Lock()
					crashedFamilies = append(crashedFamilies, family)
					mu.Unlock()
				}
			}(entryPath, entry.Name())
		}
	}

	// Wait for all processes to complete
	wg.Wait()

	// List families that crashed
	if len(crashedFamilies) > 0 {
		log.Println("The following families crashed:")
		for _, family := range crashedFamilies {
			log.Printf(" - %s", family)
		}
	} else {
		log.Println("No families crashed.")
	}
}

// runExperimentProcess runs the experiment in a separate process
func runExperimentProcess(inputPath, outputDir string) error {
	cmd := exec.Command("./experiment", "-f", inputPath, "-o", outputDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Process failed for %s: %v", inputPath, err)
		return err
	}
	log.Printf("Process completed for %s", inputPath)
	return nil
}
