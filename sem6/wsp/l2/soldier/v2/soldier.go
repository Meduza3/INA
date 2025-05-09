// Go version 2
package soldier

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const numberOfTravelers = 15
const minSteps = 10
const maxSteps = 100
const minDelay = 100 // milliseconds
const maxDelay = 500

const trapChance = 0.08

const tenantMinDelay = 100
const tenantMaxDelay = 500
const tenantLifeTime = 6000

const boardWidth = 15
const boardHeight = 15

var tenantID int64 = 15 // import "sync/atomic" at top

func nextID() int {
	return int(atomic.AddInt64(&tenantID, 1))
}

type Soldier struct {
	Id       int
	Symbol   string
	Position Position
}

type Tenant struct {
	Id       int
	Symbol   string
	Position Position
	Lifetime int
}

type EvictRequest struct {
	replyChan chan reply
}

type Position struct {
	X int
	Y int
}

type reply string

const (
	granted   reply = "granted"
	forbidden reply = "forbidden"
	trapped   reply = "trapped"
)

type TileRequest struct {
	SoldierId int
	Tenant    bool
	Enter     bool       // true for entering, false for leaving
	ReplyChan chan reply // tile replies with true if action allowed
	EvictChan chan EvictRequest
}

type TileChecker struct {
	Position
	isBusyWithSoldier bool
	isBusyWithTenant  bool
	tryChannel        chan TileRequest
	EvictChan         chan EvictRequest
	isTrap            bool
	trapId            int
}

func Start(out io.Writer) {
	printerChan := make(chan Trace, 100000)

	soldiers := make([]Soldier, 0)
	symbols := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "S"}
	for i := 0; i < numberOfTravelers; i++ {
		soldiers = append(soldiers, Soldier{Id: nextID(), Symbol: symbols[i], Position: Position{X: rand.Intn(boardWidth - 1), Y: rand.Intn(boardHeight - 1)}})
	}
	start := time.Now()
	tiles := make(map[string]*TileChecker)
	for x := 0; x < boardWidth; x++ {
		for y := 0; y < boardHeight; y++ {
			key := fmt.Sprintf("%d.%d", x, y)
			tc := &TileChecker{tryChannel: make(chan TileRequest), EvictChan: make(chan EvictRequest), Position: Position{X: x, Y: y}}
			if rand.Float64() <= trapChance {
				tc.isTrap = true
				tc.trapId = nextID()
				printerChan <- Trace{
					Time:     time.Now(),
					Id:       tc.trapId,
					Position: tc.Position,
					Symbol:   "#",
				}
			}
			tiles[key] = tc
			go tc.Run()
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(soldiers))

	go spawnTenants(printerChan, tiles, &wg)

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

func (tc *TileChecker) Run() {
	for req := range tc.tryChannel {
		// 1) Tenant is moving in or out
		if req.Tenant {
			if req.Enter {

				if tc.isTrap {
					req.ReplyChan <- trapped
					continue
				}

				// spawn or relocate: only if tile fully empty
				if !tc.isBusyWithSoldier && !tc.isBusyWithTenant {
					tc.isBusyWithTenant = true
					tc.EvictChan = req.EvictChan
					req.ReplyChan <- granted
				} else {
					req.ReplyChan <- forbidden
				}
			} else {
				// tenant life expired or final move
				tc.isBusyWithTenant = false
				tc.EvictChan = nil
				req.ReplyChan <- granted
			}
			continue
		}

		// 2) Soldier wants in or out
		if req.Enter {

			if tc.isTrap {
				req.ReplyChan <- trapped
				continue
			}

			// if there's a tenant, evict them first
			if tc.isBusyWithTenant {
				eviction := EvictRequest{replyChan: make(chan reply)}
				tc.EvictChan <- eviction
				if forbidden == <-eviction.replyChan {
					req.ReplyChan <- forbidden
					continue
				}
				// tenant has vacated
				tc.isBusyWithTenant = false
				tc.EvictChan = nil
			}
			// now try soldier
			if !tc.isBusyWithSoldier {
				tc.isBusyWithSoldier = true
				req.ReplyChan <- granted
			} else {
				req.ReplyChan <- forbidden
			}
		} else {
			tc.isBusyWithSoldier = false
			req.ReplyChan <- granted
		}
	}
}

func spawnTenants(printChan chan<- Trace, tiles map[string]*TileChecker, wg *sync.WaitGroup) {
	for {
		sleep := tenantMinDelay + rand.Intn(tenantMaxDelay-tenantMinDelay)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		t := Tenant{Symbol: fmt.Sprintf("%d", rand.Intn(9)+1)}
		go t.Run(printChan, tiles, wg)
	}
}

func (t *Tenant) Run(printChan chan<- Trace, tiles map[string]*TileChecker, wg *sync.WaitGroup) {
	t.Id = nextID()
	x := rand.Intn(boardWidth)
	y := rand.Intn(boardHeight)
	t.Position = Position{X: x, Y: y}

	key := fmt.Sprintf("%d.%d", x, y)
	replyChan := make(chan reply)
	relocateChan := make(chan EvictRequest)
	tiles[key].tryChannel <- TileRequest{
		SoldierId: t.Id,
		Enter:     true,
		Tenant:    true,
		ReplyChan: replyChan,
		EvictChan: relocateChan,
	}

	answer := <-replyChan
	if answer != granted {
		return
	}

	wg.Add(1)
	trace := Trace{
		Time:     time.Now(),
		Id:       t.Id,
		Position: t.Position,
		Symbol:   t.Symbol,
	}
	printChan <- trace

	after := time.After(tenantLifeTime * time.Millisecond)
	for {
		select {
		case <-after:
			exitChan := make(chan reply)
			tiles[key].tryChannel <- TileRequest{
				SoldierId: t.Id,
				Tenant:    true,
				Enter:     false,
				ReplyChan: exitChan,
			}
			<-exitChan
			wg.Done()
			printChan <- Trace{Time: time.Now(), Id: t.Id, Position: Position{420, 420}, Symbol: t.Symbol}
			return
		case ev := <-relocateChan:
			newX, newY := t.Position.X, t.Position.Y
		pick:
			dir := rand.Intn(4)
			switch dir {
			case 0: // Left
				newX = (t.Position.X - 1 + boardWidth) % boardWidth
			case 1: // Up
				newY = (t.Position.Y + 1) % boardHeight
			case 2: // Right
				newX = (t.Position.X + 1) % boardWidth
			case 3: // Down
				newY = (t.Position.Y - 1 + boardHeight) % boardHeight
			}
			newKey := fmt.Sprintf("%d.%d", newX, newY)
			replyChan := make(chan reply)
			tiles[newKey].tryChannel <- TileRequest{
				SoldierId: t.Id,
				Enter:     true,
				Tenant:    true,
				ReplyChan: replyChan,
				EvictChan: relocateChan,
			}
			answer := <-replyChan
			if answer == forbidden {
				goto pick
			}

			if answer == trapped {
				t.Symbol = "*"
				// Successfully entered the new tile. Now, release the previous tile.
				exitChan := make(chan reply)
				tiles[key].tryChannel <- TileRequest{
					SoldierId: t.Id,
					Enter:     false,
					ReplyChan: exitChan,
				}
				<-exitChan

				// Update soldier's position and current tile key.
				t.Position = Position{X: newX, Y: newY}
				// Send a trace of the new position.
				trace := Trace{
					Time:     time.Now(),
					Id:       t.Id,
					Position: t.Position,
					Symbol:   t.Symbol,
				}
				printChan <- trace
				time.Sleep(2 * time.Second)
				printChan <- Trace{
					Id:       tiles[key].trapId,
					Position: t.Position,
					Symbol:   "#",
					Time:     time.Now(),
				}
				return
			}
			exitChan := make(chan reply)
			tiles[key].tryChannel <- TileRequest{
				SoldierId: t.Id,
				Enter:     false,
				Tenant:    true,
				ReplyChan: exitChan,
			}
			<-exitChan
			ev.replyChan <- granted

			// Update soldier's position and current tile key.
			t.Position = Position{X: newX, Y: newY}
			key = newKey
			trace := Trace{
				Time:     time.Now(),
				Id:       t.Id,
				Position: t.Position,
				Symbol:   t.Symbol,
			}
			printChan <- trace
		}
	}
}

func (s *Soldier) Run(wg *sync.WaitGroup, printChan chan<- Trace, tiles map[string]*TileChecker) {
	defer wg.Done()
	printChan <- Trace{Time: time.Now(), Id: s.Id, Position: s.Position, Symbol: s.Symbol}
	// Determine the key for the current tile
	currentKey := fmt.Sprintf("%d.%d", s.Position.X, s.Position.Y)
	steps := minSteps + rand.Intn(maxSteps-minSteps)
	for i := 0; i < steps; i++ {
		// Wait a random delay before moving
		sleep := minDelay + rand.Intn(maxDelay-minDelay)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Decide on new position; initialize with current coordinates
	pick:
		newX, newY := s.Position.X, s.Position.Y
		dir := rand.Intn(4)
		switch dir {
		case 0: // Left
			newX = (s.Position.X - 1 + boardWidth) % boardWidth
		case 1: // Up
			newY = (s.Position.Y + 1) % boardHeight
		case 2: // Right
			newX = (s.Position.X + 1) % boardWidth
		case 3: // Down
			newY = (s.Position.Y - 1 + boardHeight) % boardHeight
		}
		newKey := fmt.Sprintf("%d.%d", newX, newY)

		// Create a reply channel for the tile entry request.
		replyChan := make(chan reply)
		tiles[newKey].tryChannel <- TileRequest{
			SoldierId: s.Id,
			Enter:     true,
			ReplyChan: replyChan,
		}

		// Wait for tile permission with a timeout to detect deadlock.
		select {
		case answer := <-replyChan:
			if answer == forbidden {
				// The tile is busy; soldier can try a different move.
				goto pick
			}
			if answer == trapped {
				s.Symbol = strings.ToLower(s.Symbol)
				// Successfully entered the new tile. Now, release the previous tile.
				exitChan := make(chan reply)
				tiles[currentKey].tryChannel <- TileRequest{
					SoldierId: s.Id,
					Enter:     false,
					ReplyChan: exitChan,
				}
				<-exitChan

				// Update soldier's position and current tile key.
				s.Position = Position{X: newX, Y: newY}
				// Send a trace of the new position.
				trace := Trace{
					Time:     time.Now(),
					Id:       s.Id,
					Position: s.Position,
					Symbol:   s.Symbol,
				}
				printChan <- trace
				time.Sleep(2 * time.Second)
				printChan <- Trace{
					Id:       tiles[fmt.Sprintf("%d.%d", s.Position.X, s.Position.Y)].trapId,
					Position: s.Position,
					Symbol:   "#",
					Time:     time.Now(),
				}
				return
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
		exitChan := make(chan reply)
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
}

type Trace struct {
	Time     time.Time
	Id       int
	Position Position
	Symbol   string
}
