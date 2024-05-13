package main

import (
	"fmt"
	"os"
)

// Program ramkujący zgodnie z zasadą "rozpychania bitów" oraz weryfikujący poprawność ramki metodą CRC.

// Odczytaj pewien źródłowy plik tekstowy 'Z' zawierający dowolny ciąg złożony ze znaków '0' i '1', symulujący strumień bitów

const (
	BEGINNING = 0
	ENDING
)

func main() {
	fmt.Println("Hello!")
	bytes, err := os.ReadFile("Z")
	if err != nil {
		fmt.Println("No file found!")
	}
	bytes_string := string(bytes)
	fmt.Println(bytes_string)
	for i, char := range bytes_string {

	}
}
