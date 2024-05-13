package main

import (
	"fmt"
	"os"
)

func main() {

	//Read file Z

	z_contents, err := os.ReadFile("z")
	check(err)

	z_contents_string := string(z_contents)
	fmt.Println(z_contents_string)

	// Open file W with write mode and create if it doesn't exist

	file, _ := os.OpenFile("w", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for _, signal_value := range z_contents_string {
		_, err := file.WriteString(fmt.Sprintf("%c", signal_value))
		check(err)
	}
	// Print to file W
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
