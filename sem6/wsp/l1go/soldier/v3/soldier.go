package soldier

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
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
	Id         int
	Symbol     string
	Position   Position
	goingUp    bool
	goingRight bool
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
		soldiers = append(soldiers, Soldier{Id: i, Symbol: symbols[i], Position: Position{X: i, Y: i}})
	}
	start := time.Now()
	tiles := make(map[string]*TileChecker)
	for x := 0; x < boardWidth; x++ {
		for y := 0; y < boardHeight; y++ {
			key := fmt.Sprintf("%d.%d", x, y)
			tc := &TileChecker{tryChannel: make(chan TileRequest)}
			tiles[key] = tc
			go tc.Run()
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(soldiers))
	for _, soldier := range soldiers {
		go soldier.Run(&wg, printerChan, tiles)
	}
	go func() {
		wg.Wait()
		close(printerChan)
	}()
	out.Write([]byte(fmt.Sprintf("-1 %d %d %d\n", numberOfTravelers, boardWidth, boardHeight)))
	for trace := range printerChan {
		out.Write([]byte(fmt.Sprintf("%f %d %d %d %s\n", time.Since(start).Seconds(), trace.Id, trace.Position.X, trace.Position.Y, trace.Symbol)))
	}
}

type TileRequest struct {
	SoldierId int
	Enter     bool      // true for entering, false for leaving
	ReplyChan chan bool // tile replies with true if action allowed
}

type TileChecker struct {
	isBusy     bool
	tryChannel chan TileRequest
}

func (tc *TileChecker) Run() {
	for req := range tc.tryChannel {
		if req.Enter {
			// Soldier is trying to enter the tile
			if !tc.isBusy {
				tc.isBusy = true
				req.ReplyChan <- true // grant access
			} else {
				req.ReplyChan <- false // deny access
			}
		} else {
			// Soldier is leaving the tile
			tc.isBusy = false
			req.ReplyChan <- true // acknowledge exit
		}
	}
}

func (s *Soldier) Run(wg *sync.WaitGroup, printChan chan<- Trace, tiles map[string]*TileChecker) {
	printChan <- Trace{Time: time.Now(), Id: s.Id, Position: s.Position, Symbol: s.Symbol}
	// Determine the direction for the run
	dir := rand.Intn(2)
	currentKey := fmt.Sprintf("%d.%d", s.Position.X, s.Position.Y)
	steps := minSteps + rand.Intn(maxSteps-minSteps)
	for i := 0; i < steps; i++ {
		// Wait a random delay before moving
		sleep := minDelay + rand.Intn(maxDelay-minDelay)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// Decide on new position; initialize with current coordinates
	pick:
		newX, newY := s.Position.X, s.Position.Y
		if s.isEven() {
			if dir == 0 {
				newY = (s.Position.Y + 1) % boardHeight
			} else {
				newY = (s.Position.Y - 1 + boardHeight) % boardHeight
			}
		} else { // s.IsOdd
			if dir == 0 {
				newX = (s.Position.X + 1) % boardWidth
			} else {
				newX = (s.Position.X - 1 + boardWidth) % boardWidth
			}
		}
		newKey := fmt.Sprintf("%d.%d", newX, newY)

		// Create a reply channel for the tile entry request.
		replyChan := make(chan bool)
		tiles[newKey].tryChannel <- TileRequest{
			SoldierId: s.Id,
			Enter:     true,
			ReplyChan: replyChan,
		}

		// Wait for tile permission with a timeout to detect deadlock.
		select {
		case granted := <-replyChan:
			if !granted {
				// The tile is busy; soldier can try a different move.
				goto pick
			}
		case <-time.After(time.Duration(maxDelay) * time.Millisecond):
			// Timeout: suspect deadlock.
			s.Symbol = strings.ToLower(s.Symbol) // Change symbol to lowercase.
			finalTrace := Trace{
				Time:     time.Now(),
				Id:       s.Id,
				Position: s.Position,
				Symbol:   s.Symbol,
			}
			printChan <- finalTrace // Leave final trace.
			wg.Done()               // Signal termination.
			return
		}

		// Successfully entered the new tile. Now, release the previous tile.
		exitChan := make(chan bool)
		tiles[currentKey].tryChannel <- TileRequest{
			SoldierId: s.Id,
			Enter:     false,
			ReplyChan: exitChan,
		}
		<-exitChan

		// Update soldier's position and current tile key.
		s.Position = Position{X: newX, Y: newY}
		currentKey = newKey

		// Send a trace of the new position.
		trace := Trace{
			Time:     time.Now(),
			Id:       s.Id,
			Position: s.Position,
			Symbol:   s.Symbol,
		}
		printChan <- trace
	}
	wg.Done()
}

func (s *Soldier) isEven() bool {
	if s.Id%2 == 0 {
		return true
	} else {
		return false
	}
}

type Trace struct {
	Time     time.Time
	Id       int
	Position Position
	Symbol   string
}
