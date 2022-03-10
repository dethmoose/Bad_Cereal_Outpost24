package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

const nChunks = 5
const chunkSize = 5

var serialPattern = regexp.MustCompile(`^(([A-Z]{5})-){4}([A-Z]{5})$`)

func validateChunk(chunk string) (int, error) {
	if len(chunk) != chunkSize {
		return 0, errors.New("chunk is of wrong size")
	}

	// Got 5 chunks, validate chunks
	sum := 0
	
	for i := 0; i < chunkSize-1; i++ {
		// Convert chunk into integer value and subtract 64 from it?
		sum += int(chunk[i]) - 64 // 'A' - 65 == 1
		// Add to sum.
	}

	// Last character in the chunk should not be equal to sum mod 26.
	// Could randomly generate a 5 character chunk and check if it has a valid checksum
	if int(chunk[chunkSize-1])-64 != sum%26 {
		return 0, fmt.Errorf("chunk checksum error")
	}
	
	// return sum
	return sum, nil
}

func validate(serial string) error {

	// Just a check if key is XXXXX-XXXXX-XXXXX-XXXXX-XXXXX
	if m := serialPattern.MatchString(serial); !m {
		return errors.New("Doesnt match pattern")
	}

	// split chunks into list
	chunks := strings.Split(serial, "-")

	// Should be 5 chunks
	if len(chunks) != nChunks {
		return errors.New("Wrong number of chunks")
	}

	// Start validation
	sum := 0

	// Run validation for first 4 chunks, adding the returned sum to sum.
	for i := 0; i < nChunks-1; i++ {
		v, err := validateChunk(chunks[i])
		if err != nil {
			return err
		}
		sum += v
	}

	// Last chunk validated.
	v, err := validateChunk(chunks[nChunks-1])
	if err != nil {
		return err
	}

	// v is sum from the last chunk
	// v != sum mod 26^4
	if int(v) != sum%int(math.Pow(26, chunkSize-1)) {
		return fmt.Errorf("serial checksum error")
	}

	return nil
}
