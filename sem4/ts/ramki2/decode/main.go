package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"strings"
)

func main() {
	// Read file W
	wContents, err := ioutil.ReadFile("w")
	if err != nil {
		fmt.Println("Error reading file W:", err)
		return
	}

	// Convert to string for processing
	wContentsString := string(wContents)

	// Remove stop sequences
	stopSequence := "01111110"
	if !strings.HasPrefix(wContentsString, stopSequence) || !strings.HasSuffix(wContentsString, stopSequence) {
		fmt.Println("Invalid framing in file W")
		return
	}
	wContentsString = strings.TrimPrefix(wContentsString, stopSequence)
	wContentsString = strings.TrimSuffix(wContentsString, stopSequence)

	// Perform bit destuffing
	destuffedSequence := destuffBits(wContentsString)

	// Separate the data and CRC
	dataLength := len(destuffedSequence) - 32
	data := destuffedSequence[:dataLength]
	crcString := destuffedSequence[dataLength:]

	// Verify CRC
	expectedCRC := generateCRC([]byte(data))
	actualCRC := uint32(0)
	for i := 0; i < 32; i++ {
		if crcString[i] == '1' {
			actualCRC |= 1 << (31 - i)
		}
	}

	if expectedCRC != actualCRC {
		fmt.Println("CRC mismatch! Data may be corrupted.")
		return
	}

	// Write the original data to file Z
	err = ioutil.WriteFile("z_recovered", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file Z:", err)
		return
	}

	fmt.Println("Data recovered successfully!")
}

// Perform bit destuffing
func destuffBits(signal string) string {
	destuffedSignal := ""
	consecutiveOnes := 0

	for i := 0; i < len(signal); i++ {
		destuffedSignal += string(signal[i])
		if signal[i] == '1' {
			consecutiveOnes++
		} else {
			consecutiveOnes = 0
		}
		if consecutiveOnes == 5 {
			// Skip the next '0'
			i++
			consecutiveOnes = 0
		}
	}

	return destuffedSignal
}

// Generate CRC32 checksum
func generateCRC(bits []byte) uint32 {
	return crc32.ChecksumIEEE(bits)
}
