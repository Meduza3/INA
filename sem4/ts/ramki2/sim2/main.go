package main

// note:
// when running this,
// redirect stderr someplace else, ex.
// go run . 2> log.txt

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
	"github.com/mattn/go-tty"
)

type signal struct {
	i int   // index
	e uint8 // owner
	b bool  // add/remove
}

var sig = make(chan signal)

type device struct {
	i int   // location
	c uint8 // char
}

var global_devices_list = make(map[uint8]device)

func eval(e map[uint8]uint) uint8 {
	var cnt = 0
	var char uint8 = 0

	mu.Lock()
	for c, v := range e {
		if v > 0 {
			cnt++
			char = c
		}
	}
	mu.Unlock()

	switch cnt {
	case 0:
		return ' '
	case 1:
		return char
	default:
		return '#'
	}
}

// convenience
func sleep(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

// length of cable
const length = 60

var cable [length]map[uint8]uint
var mu sync.Mutex

// print a char at index ind on the cable
func write(ind int, char uint8) {
	fmt.Print(cursor.RestoreAttributes() + cursor.MoveRight(ind+1))
	fmt.Printf("%c", char)
}

// handle cable state changes
func monitor(set chan device) {
	fmt.Print(cursor.Hide())
	fmt.Print(cursor.SaveAttributes())
	fmt.Println("[" + strings.Repeat(" ", length) + "]")
	for i := 0; i < length; i++ {
		cable[i] = make(map[uint8]uint)
	}

	for {
		select {
		case ev := <-sig: // event

			mu.Lock()
			if ev.b {
				cable[ev.i][ev.e] += 1
			} else {
				cable[ev.i][ev.e] -= 1
			}
			mu.Unlock()

			write(ev.i, eval(cable[ev.i]))

		case s := <-set: // add device
			fmt.Print(cursor.RestoreAttributes(),
				cursor.MoveNextLine(),
				cursor.MoveRight(s.i+1))
			fmt.Printf("%c", s.c)

			// add device to list
			global_devices_list[s.c] = s
		}
	}
}

/*

 */

func handleTTY() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		// handle key input
		var c = uint8(r)

		e, ok := global_devices_list[c]
		if ok {
			go e.CSMAsend()
		} else {
			e, ok = global_devices_list[c-32]
			if ok {
				go e.broadcast(2500, 15)
			} else {
				return
			}
		}
	}
}

/*

 */

const delay = 20 //ms

// set a tile for a set time
func (e *device) set(index, ms int, good ...bool) {
	var char uint8
	if len(good) > 0 && !good[0] {
		char = '-'
	} else {
		char = e.c
	}

	sig <- signal{index, char, true}

	go func() {
		sleep(ms)
		sig <- signal{index, char, false}
	}()
}

// send out a reverberating signal (from a device)
func (e *device) broadcast(ms, len int) {
	var rr = e.i
	var lr = e.i

	// set timer for ms milliseconds
	var timer = make(chan bool)
	go func() {
		sleep(ms)
		timer <- true
	}()

	var rr_good bool = true
	var lr_good bool = true

	var rr_end bool = false
	var lr_end bool = false

	for {
		select {
		case <-timer: // if timer has ticked, return
			return
		default:
			if eval(cable[rr]) != ' ' && eval(cable[rr]) != e.c {
				rr_good = false
			}
			if eval(cable[lr]) != ' ' && eval(cable[lr]) != e.c {
				lr_good = false
			}

			if !rr_end {
				e.set(rr, delay*len, rr_good)
				if rr < length-1 {
					rr += 1
				} else {
					rr_end = true
				}
			}

			if !lr_end {
				e.set(lr, delay*len, lr_good)
				if lr > 0 {
					lr -= 1
				} else {
					lr_end = true
				}
			}

			if rr_end && lr_end {
				return
			}

			sleep(delay)
		}
	}
}

func (e *device) CSMAsend() {
	var delay_range = 1
	for r := 0; r < 16; r++ { // repeat at most 16 times

		// wait for silence
		for eval(cable[e.i]) != ' ' {
			sleep(delay)
		}

		go e.broadcast(2*length*delay, 2*length)

		var timer = make(chan bool)
		go func() {
			sleep(2 * length * delay)
			timer <- true
		}()

		var exit = false
		var repeat = false
		for !exit {
			select {
			case <-timer:
				exit = true
			default:
				sleep(delay)
				if !repeat {
					var observed = eval(cable[e.i])
					if observed != e.c && observed != ' ' {
						// detected interference, don't recheck if already found
						log.Printf("interference detected by %c (with %c)", e.c, eval(cable[e.i]))
						repeat = true
					}
				}
			}
		}

		if repeat {
			if delay_range <= 1024 {
				delay_range = delay_range * 2
			}

			log.Printf("%c will now attempt to repeat (range %d)", e.c, delay_range)
			var off = (rand.Int()%delay_range + 1)

			log.Println(fmt.Sprintf("%c", e.c) + " picked slot " + fmt.Sprintf("%d", off))

			sleep(off * 3 * length * delay)
		} else {
			log.Println(fmt.Sprintf("%c", e.c) + " successful!")
			return
		}
	}

	log.Println(fmt.Sprintf("%c", e.c) + " failed!")
}

func main() {
	var devices = make(chan device)
	go monitor(devices)

	devices <- device{3, 'A'}
	devices <- device{20, 'B'}
	devices <- device{53, 'C'}

	handleTTY()

	// restore cursor
	fmt.Print(cursor.RestoreAttributes(),
		cursor.MoveNextLine(),
		cursor.MoveNextLine(),
		cursor.Show())
}
