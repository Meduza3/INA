package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	file, err := os.OpenFile("Z", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for range 30 {
		c := randomChar()
		charIntoBinary(file, c)
	}
	fmt.Fprintf(file, "\n")
}

func randomChar() byte {
	return byte('a' + rand.Intn(26))
}

func charIntoBinary(file *os.File, c byte) {
	fmt.Fprintf(file, "%08b", c)
}
