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
	crcString := fmt.Sprintf("%032b", crc) // CRC is 32 bits

	// Convert to string for processing
	zContentsString := string(zContents)

	// Create the sequence with CRC
	sequence := zContentsString + crcString

	// Perform bit stuffing and framing
	stopSequence := "01111110"
	framedSequence := frameSequence(sequence, stopSequence, 64)

	// Write to file W
	err = ioutil.WriteFile("w", []byte(framedSequence), 0644)
	if err != nil {
		fmt.Println("Error writing to file W:", err)
		return
	}

	fmt.Println("Framing completed successfully!")
}

func frameSequence(sequence, stopSequence string, frameSize int) string {
	finalString := ""
	for i := 0; i < len(sequence); i += frameSize {
		end := i + frameSize
		if end > len(sequence) {
			end = len(sequence)
		}
		frame := sequence[i:end]
		stuffedFrame := stuffBits(frame)
		finalString += stopSequence + stuffedFrame + stopSequence
	}
	return finalString
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
