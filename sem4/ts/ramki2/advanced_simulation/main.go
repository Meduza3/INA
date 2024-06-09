package main

import (
	"fmt"
	"sync"

	"github.com/ahmetalpbalkan/go-cursor"
)

const length = 64

var cable [length]map[uint8]uint
var mu sync.Mutex

type stacja struct {
	location int
	data     uint8
}

func monitor(stacje_channel chan stacja) {
	fmt.Print(cursor.Hide())
	fmt.Print(cursor.SaveAttributes())

}

func main() {
	var stacje = make(chan stacja)
	go monitor(stacje)
}
