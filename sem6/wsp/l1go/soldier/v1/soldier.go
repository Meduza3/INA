package soldier

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

const numberOfTravelers = 15
const minSteps = 10
const maxSteps = 100
const minDelay = 100 // milliseconds
const maxDelay = 500

const boardWidth = 15
const boardHeight = 15

type Soldier struct {
	Id       int
	Symbol   string
	Position Position
}

type Position struct {
	X int
	Y int
}

func SpawnSoldier() {

}

func Start(out io.Writer) {
	printerChan := make(chan Trace)
	soldiers := make([]Soldier, 0)
	symbols := []string{"N", "A", "D", "O", "P", "I", "E", "K", "U", "Ń", "C", "Z", "0", "Ś", "Ć"}
	for i := range numberOfTravelers {
		soldiers = append(soldiers, Soldier{Id: i, Symbol: symbols[i], Position: Position{X: rand.Intn(boardWidth), Y: rand.Intn(boardHeight)}})
	}
	start := time.Now()
	for _, soldier := range soldiers {
		go soldier.Run(printerChan)
	}
	out.Write([]byte(fmt.Sprintf("-1 %d %d %d\n", numberOfTravelers, boardWidth, boardHeight)))
	for trace := range printerChan {
		out.Write([]byte(fmt.Sprintf("%f %d %d %d %s\n", time.Since(start).Seconds(), trace.Id, trace.Position.X, trace.Position.Y, trace.Symbol)))
	}
}

func (s *Soldier) Run(printChan chan<- Trace) {
	printChan <- Trace{Time: time.Now(), Id: s.Id, Position: s.Position, Symbol: s.Symbol}
	steps := minSteps + rand.Intn(maxSteps-minSteps)
	for i := 0; i < steps; i++ {
		sleep := minDelay + rand.Intn(maxDelay-minDelay)
		blocker := time.After(time.Duration(sleep * int(time.Millisecond)))
		<-blocker
		dir := rand.Intn(4)
		switch dir {
		case 0: // Left
			s.Position.X = (s.Position.X - 1 + boardWidth) % boardWidth
		case 1: // Up
			s.Position.Y = (s.Position.Y + 1) % boardHeight
		case 2: // Right
			s.Position.X = (s.Position.X + 1) % boardWidth
		case 3: // Down
			s.Position.Y = (s.Position.Y - 1 + boardHeight) % boardHeight
		}
		printChan <- Trace{Time: time.Now(), Id: s.Id, Position: s.Position, Symbol: s.Symbol}

	}
}

type Trace struct {
	Time     time.Time
	Id       int
	Position Position
	Symbol   string
}
