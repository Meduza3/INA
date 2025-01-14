package graphs

func RunSelectedAlgorithm(g *Graph, source int, algoName string) ([]int, error) {
	switch algoName {
	case "dijkstra":
		dist, _, err := DijkstraBasic(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	case "dial":
		dist, _, err := DijkstraDial(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	case "radixheap":
		dist, _, err := DijkstraRadix(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	default:
		dist, _, err := DijkstraBasic(g, source)
		if err != nil {
			return nil, err
		}
		return dist, nil
	}
}
