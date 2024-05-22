package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var widelce = [ilosc_filozofow]Widelec{}
	var filozofowie = [ilosc_filozofow]*Filozof{}
	var waitGroup = sync.WaitGroup{}

	for i := 0; i < ilosc_filozofow; i++ {
		var lewa = &widelce[i]
		var prawa = &widelce[(i+1)%ilosc_filozofow]

		var filozof *Filozof
		if i%2 == 0 {
			filozof = nowyFilozof(i, prawa, lewa, ilosc_posilkow, &waitGroup)
		} else {
			filozof = nowyFilozof(i, lewa, prawa, ilosc_posilkow, &waitGroup)
		}
		filozofowie[i] = filozof
	}

	for _, f := range filozofowie {
		waitGroup.Add(1)
		go f.start()
	}
	waitGroup.Wait()
}

const (
	ilosc_filozofow int = 5

	min_dobranoc   int           = 50
	max_dobranoc   int           = 100
	jeden_dzien    time.Duration = time.Millisecond
	ilosc_posilkow int           = 5
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Filozof struct {
	id             int
	left           *Widelec
	right          *Widelec
	ilosc_posilkow int
	waitGroup      *sync.WaitGroup
}

func nowyFilozof(id int, left *Widelec, right *Widelec, ilosc_posilkow int, waitGroup *sync.WaitGroup) *Filozof {
	return &Filozof{id, left, right, ilosc_posilkow, waitGroup}
}

func (f *Filozof) mysl() {
	fmt.Printf("Filozof %d mysli \n", f.id)
	f.dobranoc()
	fmt.Printf("Filozof %d skonczyl myslec \n", f.id)
}

func (f *Filozof) jedz() {
	fmt.Printf("Filozof %d je \n", f.id)
	f.dobranoc()
	fmt.Printf("Filozof %d skonczyl jesc \n", f.id)
}

func (f *Filozof) podnies_widelec(widelec *Widelec) {
	f.dobranoc()
	widelec.Lock()
}

func (f *Filozof) odloz_widelec(widelec *Widelec) {
	f.dobranoc()
	widelec.Unlock()
}

func (f *Filozof) dobranoc() {
	time.Sleep(time.Duration(rng.Intn(min_dobranoc)+max_dobranoc) * jeden_dzien)
}

func (f *Filozof) start() {
	defer f.waitGroup.Done()
	for i := 0; i < f.ilosc_posilkow; i++ {
		f.mysl()
		f.podnies_widelec(f.left)
		f.podnies_widelec(f.right)
		f.jedz()
		f.odloz_widelec(f.left)
		f.odloz_widelec(f.right)
	}
}

type Widelec struct {
	sync.Mutex
}
