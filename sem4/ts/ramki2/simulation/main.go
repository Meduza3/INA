package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	STACJA_COUNT  = 10
	MEDIUM_LENGTH = 64
	DATA_LENGTH   = 128
)

var (
	stacje []Stacja
	kabel  Kabel
)

func main() {

	kabel = make(Kabel, MEDIUM_LENGTH)
	for i := range kabel {
		kabel[i] = make([]int, STACJA_COUNT)
		for j := range kabel[i] {
			kabel[i][j] = -1
		}
	}
	// Randomize stacja placements

	usedLocations := make(map[int]bool)
	for i := 0; i < STACJA_COUNT; i++ {
		var location int
		for {
			location = rand.Intn(MEDIUM_LENGTH)
			if !usedLocations[location] {
				break
			}
		}
		usedLocations[location] = true
		stacja := Stacja{
			id:        i,
			location:  location,
			data_left: DATA_LENGTH,
			my_place:  &kabel[location][i],
			jam_left:  DATA_LENGTH,
		}
		stacje = append(stacje, stacja)
	}

	for {
		for _, stacja := range stacje {
			stacja.nasluchuj()
		}
		kabel.propaguj()

		fmt.Print("\033[2J") // 100 newline characters
		kabel.print_simple_kabel()
		kabel.print_kabel()
		time.Sleep(time.Millisecond * 500)
	}

}

type Kabel [][]int

func (k Kabel) print_kabel() {
	for i := 0; i < STACJA_COUNT; i++ {
		fmt.Printf("%d: ", i)
		for j := 0; j < MEDIUM_LENGTH; j++ {
			var atLocation = false
			if j == stacje[i].location {
				fmt.Print("\033[44m")
				atLocation = true
			}
			switch kabel[j][i] {
			case -1:
				fmt.Print("-")
			case -99:
				if atLocation {
					fmt.Print("X")
				} else {
					fmt.Print("\033[41mX")
				}
			default:
				fmt.Print(kabel[j][i])
			}
			fmt.Print("\033[0m")
		}
		fmt.Println()
	}
}

func (k Kabel) print_simple_kabel() {
	fmt.Printf("K: ")
	simple_kabel := make([]int, MEDIUM_LENGTH)
	for i := 0; i < MEDIUM_LENGTH; i++ {
		simple_kabel[i] = -1
	}

	for i := 0; i < STACJA_COUNT; i++ {
		for j := 0; j < MEDIUM_LENGTH; j++ {
			if k[j][i] != -1 {
				simple_kabel[j] = k[j][i]
			}
		}
	}

	for miejsce := 0; miejsce < MEDIUM_LENGTH; miejsce++ {
		switch simple_kabel[miejsce] {
		case -1:
			fmt.Print("-")
		case -99:
			fmt.Print("\033[41mX\033[0m")
		default:
			fmt.Print(simple_kabel[miejsce])
		}
	}
	fmt.Println()
}

func (k *Kabel) propaguj() {
	for i := 0; i < STACJA_COUNT; i++ {
		j := 0
		for j < MEDIUM_LENGTH {
			if (*k)[j][i] != -1 {
				propagated := false
				if j < stacje[i].location {
					if j > 0 && (*k)[j-1][i] != -99 {
						(*k)[j-1][i] = (*k)[j][i]
						propagated = true
					}
				}
				if j >= stacje[i].location && j < MEDIUM_LENGTH-1 && (*k)[j+1][i] != -99 {
					(*k)[j+1][i] = (*k)[j][i]
					propagated = true
				}

				if propagated {
					j += 2
				} else {
					j++
				}
			} else {
				j++
			}
		}
	}
}

type Stacja struct {
	id        int
	location  int
	data_left int
	my_place  *int
	jam_left  int
}

func (s *Stacja) nasluchuj() {
	collision := false
	for i := 0; i < STACJA_COUNT; i++ {
		if kabel[s.location][i] != -1 && kabel[s.location][i] != s.id {
			collision = true
		}
	}
	if s.data_left > 0 {
		if collision {
			*s.my_place = -99
			s.jam_left--

		} else {
			*s.my_place = s.id
			s.data_left--
		}
	}
}
