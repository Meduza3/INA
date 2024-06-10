package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

type Node struct {
	name               string
	position           int
	collision          bool // Did a collision happen?
	broadcasting       bool // Is it broadcasting right now?
	wait               int  // Number of iterations node will wait for
	left               int  // Number of frames left to broadcast
	total_collisions   int
	total_waiting_time int
}

func newNode(name string, position int, frame_count int) *Node {
	return &Node{name: name, position: position, left: frame_count}
}

func (n *Node) flag() string {
	if n.broadcasting {
		if n.collision {
			return "Collided"
		} else {
			return "Active"
		}
	} else {
		if n.left == 0 {
			return "Idle"
		} else {
			return "Waiting"
		}
	}
}

func (n *Node) tryLater() {
	n.total_collisions += 1
	n.wait = rand.Intn(int(math.Pow(2, math.Min(float64(n.total_collisions), 10))))
	n.total_waiting_time += n.wait
}

type Signal struct {
	node  *Node
	isJam bool
}

const (
	CELL_WIDTH = 10
)

type Simulation struct {
	nodes        []*Node
	active_nodes []*Node
	cable        [][]*Signal
	empty_cable  [][]*Signal
	size         int
	log          []string
	names        []string
}

func newSimulation(size int) *Simulation {
	nodes := make([]*Node, size)
	for i := range nodes {
		nodes[i] = &Node{}
	}
	names := make([]string, size)
	cable := make([][]*Signal, size)
	for i := range cable {
		cable[i] = make([]*Signal, size)
	}

	empty_cable := make([][]*Signal, size)
	for i := range cable {
		empty_cable[i] = make([]*Signal, size)
	}

	return &Simulation{
		nodes:        nodes,
		active_nodes: make([]*Node, 0),
		cable:        cable,
		empty_cable:  empty_cable,
		size:         size,
		names:        names,
	}
}

func (s *Simulation) add_node(name string, position int, frames_to_send int) error {
	if position < 0 || position >= s.size {
		return fmt.Errorf("Invalid position")
	}
	node := newNode(name, position, frames_to_send)
	s.nodes[position] = node
	s.names[position] = name

	if frames_to_send > 0 {
		s.active_nodes = append(s.active_nodes, node)
	}

	return nil

}

func (s *Simulation) run() {
	i := -1
	reader := bufio.NewReader(os.Stdin)
	for len(s.active_nodes) > 0 || !reflect.DeepEqual(s.cable, s.empty_cable) {
		i += 1
		s.step()
		s.printDebugState(i)
		fmt.Println("Stepping!")
		reader.ReadString('\n')
	}

}

func (s *Simulation) step() {
	next_cable := deepCopyCable(s.empty_cable)

	// First, propagate existing signals
	for i, segment := range s.cable {
		for _, signal := range segment {
			if i > 0 { // Propagate left
				next_cable[i-1] = append(next_cable[i-1], signal)
			}
			if i < s.size-1 { // Propagate right
				next_cable[i+1] = append(next_cable[i+1], signal)
			}
		}
	}

	// Check for collisions at each position
	for _, signals := range next_cable {
		if len(signals) > 1 {
			// Collision detected, mark all involved nodes
			for _, signal := range signals {
				signal.node.collision = true
			}
		}
	}

	// Handle broadcasting attempts
	for _, node := range s.active_nodes {
		if node.wait == 0 && !node.collision {
			if len(s.cable[node.position]) == 0 {
				node.broadcasting = true
				next_cable[node.position] = append(next_cable[node.position], &Signal{node: node, isJam: false})
			}
		}

		if node.broadcasting {
			node.left--
			if node.left == 0 || node.collision {
				node.broadcasting = false
			}
		}

		if node.collision {
			node.tryLater()        // Apply backoff after collision
			node.collision = false // Reset collision state
		}
	}

	s.cable = next_cable
}

func (s *Simulation) removeActiveNode(node *Node) {
	for i, activeNode := range s.active_nodes {
		if activeNode == node {
			s.active_nodes = append(s.active_nodes[:i], s.active_nodes[i+1:]...)
			return
		}
	}
}

func deepCopyCable(original [][]*Signal) [][]*Signal {
	copy := make([][]*Signal, len(original))
	for i := range original {
		copy[i] = make([]*Signal, 0) // Ensure slices are initialized empty
		for _, signal := range original[i] {
			if signal != nil {
				// Assuming Signal struct contains fields that should be deeply copied
				newSignal := &Signal{
					node:  signal.node, // Point to the same node, typically not deep-copied
					isJam: signal.isJam,
				}
				copy[i] = append(copy[i], newSignal)
			}
		}
	}
	return copy
}

func (s *Simulation) printHeader() {
	w := CELL_WIDTH
	fmt.Println("+" + strings.Repeat("-", 6) + "+" + strings.Repeat("-"+strings.Repeat(" ", w), s.size) + "+")
	fmt.Printf("|%6s|", "Name")
	for _, name := range s.names {
		fmt.Printf("%*s|", w, name)
	}
	fmt.Println()
	fmt.Println("+" + strings.Repeat("=", 6) + "+" + strings.Repeat("="+strings.Repeat(" ", w), s.size) + "+")
	fmt.Printf("|%6s|", "Wait")
	for _, node := range s.nodes {
		if node != nil {
			fmt.Printf("%*d|", w, node.wait)
		} else {
			fmt.Printf("%*s|", w, "")
		}
	}
	fmt.Println()
	fmt.Printf("|%6s|", "Flag")
	for _, node := range s.nodes {
		if node != nil {
			fmt.Printf("%*s|", w, node.flag())
		} else {
			fmt.Printf("%*s|", w, "")
		}
	}
	fmt.Println()
	fmt.Println("+" + strings.Repeat("=", 6) + "+" + strings.Repeat("="+strings.Repeat(" ", w), s.size) + "+")
}

func (s *Simulation) printCable(iter int) {
	w := CELL_WIDTH
	fmt.Printf("|%6d|", iter)
	for _, segment := range s.cable {
		if segment != nil {
			names := []string{}
			for _, signal := range segment {
				if signal != nil && signal.node != nil {
					names = append(names, signal.node.name)
				}
			}
			if len(names) > 0 {
				fmt.Printf("%*s|", w, strings.Join(names, ", "))
			} else {
				fmt.Printf("%*s|", w, "")
			}
		}
	}
	fmt.Println()
	fmt.Println("+" + strings.Repeat("-", 6) + "+" + strings.Repeat("-"+strings.Repeat(" ", w), s.size) + "+")
}

func (s *Simulation) printState(iter int) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	s.printHeader()
	s.printCable(iter)

}

func (s *Simulation) printDebugState(iter int) {
	fmt.Printf("Iteration: %d\n", iter)
	fmt.Println("Cable State:")
	for i, segment := range s.cable {
		if len(segment) > 0 {
			fmt.Printf("Position %d: ", i)
			for _, signal := range segment {
				if signal != nil && signal.node != nil {
					fmt.Printf("%s ", signal.node.name)
				}
			}
			fmt.Println()
		}
	}
	fmt.Println("Node States:")
	for _, node := range s.nodes {
		if node != nil {
			fmt.Printf("%s - Wait: %d, Broadcasting: %t, Collision: %t\n", node.name, node.wait, node.broadcasting, node.collision)
		}
	}
	fmt.Println("-----------------------------------")
}

func main() {
	sim := newSimulation(10)
	sim.add_node("A", 0, 1)
	sim.add_node("B", 2, 1)
	sim.add_node("C", 7, 1)
	sim.run()
}
