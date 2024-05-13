package main

import (
	"fmt"
	"os"
)

func main() {

	//Read file Z

	z_contents, _ := os.ReadFile("z")

	z_contents_string := string(z_contents)
	fmt.Println(z_contents_string)
	header := "_____header_____"
	crc := "______crc_______"

	final_string := "01111110"
	final_string += header
	//Bit stuffing - Dodaj 0 po pięciu jedynkach na raz
	final_string += stuff_bits(z_contents_string)
	final_string += crc
	final_string += "01111110"
	file, _ := os.OpenFile("w", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	_, err := file.WriteString(final_string)
	check(err)
}

// Print to file W

func stuff_bits(signal string) string {
	stuffed_signal := ""
	consecutive_ones := 0
	for _, signal_value := range signal {
		stuffed_signal = stuffed_signal + string(signal_value)
		if string(signal_value) == "1" {
			consecutive_ones++
		}
		if consecutive_ones == 5 {
			consecutive_ones = 0
			stuffed_signal = stuffed_signal + "0"
		}
	}
	return stuffed_signal
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
