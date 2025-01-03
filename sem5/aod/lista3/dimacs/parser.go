package dimacs

import (
	"dijkstra/graphs"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParsePairs(f *os.File) ([]struct{ From, To int }, error) {
	text, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(text), "\n")
	var pairs []struct{ From, To int }

	for _, line := range lines {
		// Skip empty lines or comment lines
		if len(line) == 0 || line[0] == 'c' {
			continue
		}
		// Skip the problem definition line
		if line[0] == 'p' {
			continue
		}
		// Parse lines starting with 'q'
		if line[0] == 'q' {
			fields := strings.Fields(line)
			if len(fields) < 3 {
				return nil, fmt.Errorf("invalid pair definition line: %s", line)
			}
			from, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, fmt.Errorf("invalid 'from' vertex in line %q: %w", line, err)
			}
			to, err := strconv.Atoi(fields[2])
			if err != nil {
				return nil, fmt.Errorf("invalid 'to' vertex in line %q: %w", line, err)
			}
			pairs = append(pairs, struct{ From, To int }{
				From: from,
				To:   to,
			})
		}
	}

	return pairs, nil
}

func ParseSources(f *os.File) ([]int, error) {
	text, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	lines := strings.Split(string(text), "\n")
	var sources []int
	for _, line := range lines {
		if len(line) == 0 || line[0] == 'c' || line[0] == 'p' {
			continue
		}
		if line[0] == 's' {
			fields := strings.Fields(line)
			source, err := strconv.Atoi(fields[1])
			if err != nil {
				continue
			}
			sources = append(sources, source)
		}
	}
	return sources, nil
}

func GenerateSSResults(graphFile, sourceFile string, time time.Duration, size, edgeCount, maxCost, minCost int) string {
	return fmt.Sprintf(`
															f %s %s
															g %d %d %d %d
															t %.2f
															`,
		graphFile, sourceFile,
		size, edgeCount, minCost, maxCost,
		time.Seconds()*1000)
}

func GenerateP2PResults(graphFile, pairFile string, size, edgeCount, maxCost, minCost int, distances []struct {
	From, To, Distance int
}) string {
	result := fmt.Sprintf(`
f %s %s
g %d %d %d %d
`,
		graphFile, pairFile,
		size, edgeCount, minCost, maxCost)

	for _, d := range distances {
		result += fmt.Sprintf("d %d %d %d\n", d.From, d.To, d.Distance)
	}

	return result
}

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
