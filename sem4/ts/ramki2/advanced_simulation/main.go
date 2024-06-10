package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Node struct {
	name             string
	position         int
	collision        bool
	collisionCount   int
	broadcasting     bool
	wait             int
	left             int
	totalCollisions  int
	totalWaitingTime int
}

func (n *Node) flag() string {
	if n.broadcasting {
		if n.collision {
			return "Collided"
		}
		return "Active"
	}
	if n.left == 0 {
		return "Idle"
	}
	return "Waiting"
}

func (n *Node) reschedule() {
	n.collisionCount++
	n.wait = rand.Intn(1 << min(n.collisionCount, 10))
	n.totalWaitingTime += n.wait
}

type Signal struct {
	node      *Node
	direction int
	jam       bool
}

func (s *Signal) String() string {
	if s.jam {
		return string(s.node.name[0]) + "*"
	}
	return string(s.node.name[0])
}

type Simulation struct {
	nodes       map[int]*Node
	activeNodes []*Node
	cable       [][]*Signal
	emptyCable  [][]*Signal
	size        int
	names       []string
	width       int
	log         []string
}

func NewSimulation(size int) *Simulation {
	sim := &Simulation{
		nodes:       make(map[int]*Node, size),
		activeNodes: []*Node{},
		cable:       make([][]*Signal, size),
		emptyCable:  make([][]*Signal, size),
		size:        size,
		names:       make([]string, size),
		width:       10,
		log:         []string{},
	}
	for i := range sim.cable {
		sim.cable[i] = []*Signal{}
		sim.emptyCable[i] = []*Signal{}
	}
	return sim
}

func (sim *Simulation) AddNode(name string, position int, start int, framesToSend int) {
	node := &Node{name: name, position: position, wait: start, left: framesToSend}
	if _, exists := sim.nodes[position]; !exists && position >= 0 && position < sim.size {
		sim.nodes[position] = node
		sim.names[position] = name
		if framesToSend > 0 {
			sim.activeNodes = append(sim.activeNodes, node)
		}
	} else {
		panic(fmt.Sprintf("Position must be between 0 and %d.", sim.size-1))
	}
}

func (sim *Simulation) Run(outputAll bool, displayTime int) {
	sim.width = max(4*len(sim.activeNodes)-2, 10)
	if outputAll {
		sim.printHeader(true)
	}
	i := -1
	for len(sim.activeNodes) > 0 || !sim.isEmptyCable() {
		i++
		sim.step()
		if outputAll {
			sim.printCable(i)
		} else {
			sim.printState(i)
			if displayTime > 0 {
				time.Sleep(time.Duration(displayTime) * time.Second)
			} else {
				fmt.Print("Press Enter to continue...")
				fmt.Scanln()
			}
		}
	}
	fmt.Printf("\n%d iterations total. Node statistics:\n", i)
	for _, node := range sim.nodes {
		if node != nil {
			fmt.Printf("%s\ncollisions: %d\ntotal waiting time: %d\n\n",
				node.name, node.totalCollisions, node.totalWaitingTime)
		}
	}
}

func (sim *Simulation) step() {
	// Create a new instance of the cable for the next state
	nextCable := make([][]*Signal, sim.size)
	for i := range nextCable {
		nextCable[i] = []*Signal{}
	}

	// Propagate signals in the cable
	for i, segment := range sim.cable {
		for _, signal := range segment {
			if signal.direction == -1 {
				if i > 0 {
					nextCable[i-1] = append(nextCable[i-1], signal)
				}
			} else if signal.direction == 1 {
				if i < sim.size-1 {
					nextCable[i+1] = append(nextCable[i+1], signal)
				}
			} else {
				if i > 0 {
					nextCable[i-1] = append(nextCable[i-1], &Signal{node: signal.node, direction: -1, jam: signal.jam})
				}
				if i < sim.size-1 {
					nextCable[i+1] = append(nextCable[i+1], &Signal{node: signal.node, direction: 1, jam: signal.jam})
				}
			}
		}
	}

	sim.cable = nextCable

	// Process each node's state and interaction with the cable
	for _, node := range sim.activeNodes {
		if !node.broadcasting {
			if node.wait == 0 {
				if len(sim.cable[node.position]) == 0 {
					sim.log = append(sim.log, fmt.Sprintf("%s started broadcasting.", node.name))
					sim.cable[node.position] = append(sim.cable[node.position], &Signal{node: node, direction: 0, jam: false})
					node.broadcasting = true
					node.wait = 2*sim.size - 2
				} else {
					node.totalWaitingTime++
				}
			} else {
				node.wait--
			}
		} else {
			if node.wait == 0 {
				node.broadcasting = false
				if node.collision {
					node.collision = false
					node.reschedule()
					sim.log = append(sim.log, fmt.Sprintf("%s is done jamming. It will wait %d iteration(s) before trying to broadcast again.", node.name, node.wait))
				} else {
					node.left--
					message := "Its work is done."
					if node.left > 0 {
						message = fmt.Sprintf("%d to go.", node.left)
					}
					sim.log = append(sim.log, fmt.Sprintf("%s has successfully broadcasted one of its frames. %s", node.name, message))
					node.totalCollisions += node.collisionCount
					node.collisionCount = 0
					if node.left == 0 {
						sim.activeNodes = removeNode(sim.activeNodes, node)
					}
				}
			} else {
				if !node.collision && len(sim.cable[node.position]) > 1 {
					sim.log = append(sim.log, fmt.Sprintf("%s has detected a collision. It started to broadcast its jam signal.", node.name))
					node.collision = true
					node.wait = 2*sim.size - 2
				}
				sim.log = append(sim.log, fmt.Sprintf("%s continues to broadcast.", node.name))
				sim.cable[node.position] = append(sim.cable[node.position], &Signal{node: node, direction: 0, jam: node.collision})
				node.wait--
			}
		}
	}
}

func removeNode(slice []*Node, node *Node) []*Node {
	for i, n := range slice {
		if n == node {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (sim *Simulation) isEmptyCable() bool {
	for _, segment := range sim.cable {
		if len(segment) > 0 {
			return false
		}
	}
	return true
}

func (sim *Simulation) printHeader(universal bool) {
	width := sim.width
	fmt.Printf("+%s+\n", strings.Repeat("-", 6+width*sim.size))
	fmt.Printf("|%-6s|", "Name")
	for _, name := range sim.names {
		fmt.Printf("%-*s|", width, name)
	}
	fmt.Println()

	if !universal {
		fmt.Printf("|%-6s|", "Wait")
		for _, node := range sim.nodes {
			if node != nil {
				fmt.Printf("%-*d|", width, node.wait)
			} else {
				fmt.Printf("%-*s|", width, "")
			}
		}
		fmt.Println()

		fmt.Printf("|%-6s|", "Flag")
		for _, node := range sim.nodes {
			if node != nil {
				fmt.Printf("%-*s|", width, node.flag())
			} else {
				fmt.Printf("%-*s|", width, "")
			}
		}
		fmt.Println()
	}

	fmt.Printf("+%s+\n", strings.Repeat("=", 6+width*sim.size))
}

func (sim *Simulation) printCable(iter int) {
	width := sim.width
	fmt.Printf("+%s+\n", strings.Repeat("-", 6+width*sim.size))
	fmt.Printf("|%-6s|", fmt.Sprintf("t=%d", iter))
	for _, segment := range sim.cable {
		var signals []string
		for _, signal := range segment {
			signals = append(signals, signal.String())
		}
		fmt.Printf("%-*s|", width, strings.Join(signals, ", "))
	}
	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 6+width*sim.size))
}

func (sim *Simulation) printState(iter int) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // Ensure you've imported "os/exec"
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J") // Terminal clear command for Unix-like systems
	}
	sim.printHeader(false)
	sim.printCable(iter)
	for _, message := range sim.log {
		fmt.Println(message)
	}
	sim.log = []
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	sim := NewSimulation(5)
	sim.AddNode("A", 0, 0, 1)
	sim.AddNode("B", 2, 0, 0)
	sim.AddNode("C", 4, 3, 1)
	sim.Run(true, 0)
}
