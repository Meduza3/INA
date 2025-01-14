package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const baseDir = "tests"

var folders []string

func main() {

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			relPath, err := filepath.Rel(baseDir, path)
			if err != nil {
				return err
			}
			folders = append(folders, relPath)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to get folders in %s:%v", baseDir, err)
	}

	type Result struct {
		folder string
		time   time.Duration
	}
	results := make([]Result, 0)

	for _, folder := range folders {
		fmt.Printf("in %s\n", folder)
		cmd := exec.Command("julia", "main.jl")
		in, err := cmd.StdinPipe()
		if err != nil {
			log.Fatalf("Failed to get input pipe: %v", err)
		}
		out, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatalf("Failed to get output pipe: %v", err)
		}
		now := time.Now()
		err = cmd.Start()
		if err != nil {
			log.Fatalf("Failed to get run command: %v", err)
		}
		input := fmt.Sprintf("experiment tests/%s/A.txt tests/%s/b.txt", folder, folder)
		in.Write([]byte(input))
		in.Close()
		output, err := io.ReadAll(out)
		fmt.Println(string(output[:100]))
		res := Result{
			folder: folder,
			time:   time.Since(now),
		}
		fmt.Println(res)
		results = append(results, res)
	}

	fmt.Println("REZULTATY: ")
	for _, res := range results {
		fmt.Println(res)

	}

}
