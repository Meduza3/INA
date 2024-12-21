package dimacs

import (
	"dijkstra/graphs"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ParseFile(f *os.File) (*graphs.Graph, error) {
	text, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var g *graphs.Graph
	lines := strings.Split(string(text), "\n")
	for _, line := range lines {
		// Skip empty lines
		if len(line) == 0 {
			continue
		}
		// Skip comment lines starting with 'c'
		if line[0] == 'c' {
			continue
		}
		if line[0] == 'p' {
			fields := strings.Fields(line)
			if len(fields) < 4 || fields[1] != "sp" {
				return nil, fmt.Errorf("invalid problem definition line: %s", line)
			}
			numVertices, err := strconv.Atoi(fields[2])
			if err != nil {
				return nil, fmt.Errorf("invalid number of vertices: %w", err)
			}
			g = graphs.NewGraph(numVertices)
		}
		if line[0] == 'a' {
			fields := strings.Fields(line)
			if len(fields) < 4 {
				return nil, fmt.Errorf("invalid arc definition line: %s", line)
			}
			from, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, fmt.Errorf("invalid 'from' vertex: %w", err)
			}
			to, err := strconv.Atoi(fields[2])
			if err != nil {
				return nil, fmt.Errorf("invalid 'to' vertex: %w", err)
			}
			cost, err := strconv.Atoi(fields[3])
			if err != nil {
				return nil, fmt.Errorf("invalid cost: %w", err)
			}
			// Adjust for 1-based indexing in the DIMACS format
			err = g.AddEdge(from-1, to-1, cost)
			if err != nil {
				return nil, fmt.Errorf("failed to add edge: %w", err)
			}
		}
		if g == nil {
			return nil, fmt.Errorf("no graph definition found in the file")
		}
	}
	return g, nil
}
