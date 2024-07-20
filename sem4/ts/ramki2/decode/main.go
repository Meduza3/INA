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

	// Split into frames by the stop sequence
	frames := strings.Split(wContentsString, "01111110")
	var allData string

	for _, frame := range frames {
		if frame == "" {
			continue
		}

		// Perform bit destuffing on each frame
		destuffed := destuffBits(frame)
		allData += destuffed
	}

	if len(allData) < 32 {
		fmt.Println("Invalid or insufficient data for CRC check.")
		return
	}

	// Separate the data and CRC
	dataLength := len(allData) - 32
	data := allData[:dataLength]
	crcString := allData[dataLength:]

	// Convert binary string CRC to uint32
	actualCRC := uint32(0)
	for i := 0; i < 32; i++ {
		if crcString[i] == '1' {
			actualCRC |= 1 << (31 - i)
		}
	}

	// Verify CRC
	expectedCRC := generateCRC([]byte(data))
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
			if i+1 < len(signal) && signal[i+1] == '0' {
				i++ // Skip the '0' after five consecutive '1's
			}
			consecutiveOnes = 0
		}
	}

	return destuffedSignal
}

// Generate CRC32 checksum
func generateCRC(bits []byte) uint32 {
	return crc32.ChecksumIEEE(bits)
}
