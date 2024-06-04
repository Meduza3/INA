package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func main() {
	// Read file Z
	zContents, err := ioutil.ReadFile("z")
	if err != nil {
		fmt.Println("Error reading file Z:", err)
		return
	}

	// Calculate CRC
	crc := generateCRC(zContents)
	crcString := fmt.Sprintf("%032b", crc)

	// Convert to string for processing
	zContentsString := string(zContents)

	// Create the sequence with CRC
	sequence := zContentsString + crcString

	// Perform bit stuffing
	stuffedSequence := stuffBits(sequence)

	// Frame the sequence
	stopSequence := "01111110"
	finalString := stopSequence + stuffedSequence + stopSequence

	// Write to file W
	err = ioutil.WriteFile("w", []byte(finalString), 0644)
	if err != nil {
		fmt.Println("Error writing to file W:", err)
		return
	}

	fmt.Println("Framing completed successfully!")
}

// Perform bit stuffing
func stuffBits(signal string) string {
	stuffedSignal := ""
	consecutiveOnes := 0

	for _, signalValue := range signal {
		stuffedSignal += string(signalValue)
		if signalValue == '1' {
			consecutiveOnes++
		} else {
			consecutiveOnes = 0
		}
		if consecutiveOnes == 5 {
			stuffedSignal += "0"
			consecutiveOnes = 0
		}
	}

	return stuffedSignal
}

// Generate CRC32 checksum
func generateCRC(bits []byte) uint32 {
	return crc32.ChecksumIEEE(bits)
}
